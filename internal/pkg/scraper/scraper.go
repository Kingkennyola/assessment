// Package handlers implements handler functions used to handle
// requests in the service.
package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"phaidra-assessment/internal/pkg/config"
	"phaidra-assessment/internal/pkg/metrics"
	"phaidra-assessment/internal/pkg/models"

	"github.com/prometheus/client_golang/prometheus"
)

func Scraper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Empty request body.")
			return
		}

		var scraperRequest models.ScraperRequest

		err := decoder.Decode(&scraperRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		if IsUrl(scraperRequest.URL) {
			w.WriteHeader(http.StatusCreated)
			go makeRequest(scraperRequest.URL)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Invalid URL in request body.")
			return
		}
	}
}

func makeRequest(url string) {
	client := http.Client{
		Timeout: config.NewConfig().ScraperRequestTimeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		log.Println(err)
	} else {
		metrics.ScraperRequestCounter.With(prometheus.Labels{"url": url, "code": fmt.Sprint(resp.StatusCode)}).Inc()
	}
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		log.Println(err)
	}
	return err == nil && u.Scheme != "" && u.Host != ""
}
