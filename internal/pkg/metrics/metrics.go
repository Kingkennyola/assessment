package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetricsHandler() http.Handler {
	reg := prometheus.NewRegistry()
	reg.MustRegister(ScraperRequestCounter)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})
	return promHandler
}

var ScraperRequestCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_get",
		Help: "HTTP GET requests",
	}, []string{"url", "code"},
)
