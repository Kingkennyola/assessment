package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phaidra-assessment/internal/pkg/config"
	"phaidra-assessment/internal/pkg/health"
	"phaidra-assessment/internal/pkg/metrics"
	"phaidra-assessment/internal/pkg/scraper"
	"time"
)

var cfg *config.Config

func main() {
	cfg = config.NewConfig()

	r := http.NewServeMux()
	r.HandleFunc("POST /", scraper.Scraper())
	r.HandleFunc("GET /health", health.HealthGet())

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ListenPort),
		Handler:      r,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
		IdleTimeout:  cfg.ServerIdleTimeout,
	}

	m := http.NewServeMux()
	mh := metrics.NewMetricsHandler()
	m.Handle("GET /metrics", mh)

	ms := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.MetricsPort),
		Handler:      m,
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
		IdleTimeout:  cfg.ServerIdleTimeout,
	}

	log.Printf("Server is running on port %s\n", cfg.ListenPort)
	go s.ListenAndServe()

	log.Printf("Metrics is running on port %s\n", cfg.MetricsPort)
	go ms.ListenAndServe()

	// channel for ctrl+c/SIGINT event
	sigInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(sigInterruptChannel, os.Interrupt)

	// block execution until SIGINT comes
	<-sigInterruptChannel

	// context with 4 seconds grace period
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	// shutdown both servers
	go s.Shutdown(ctx)
	go ms.Shutdown(ctx)

	// wait for ctx end
	<-ctx.Done()
}
