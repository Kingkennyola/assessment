// Tests for the handlers handlers package.
package scraper

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"phaidra-assessment/internal/pkg/health"
	"phaidra-assessment/internal/pkg/metrics"

	"github.com/stretchr/testify/assert"
)

func TestScraper(t *testing.T) {
	tests := map[string]struct {
		body   []byte
		method string
		result int
	}{
		"CorrectURL": {
			body:   []byte(`{"url":"http://ifconfig.me"}`),
			method: http.MethodPost,
			result: http.StatusCreated,
		},
		"BadURL": {
			body:   []byte(`{"url":"http:/ifconfig.me"}`),
			method: http.MethodPost,
			result: http.StatusBadRequest,
		},
		"BadJSON": {
			body:   []byte(`{"a":"b"}`),
			method: http.MethodGet,
			result: http.StatusBadRequest,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(fmt.Sprintf("TestScraperWith%s", name), func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(test.method, "/", bytes.NewBuffer(test.body))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Scraper())
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.result, rr.Code)
		})
	}
}

func TestMetrics(t *testing.T) {
	tests := map[string]struct {
		url      string
		response int
	}{
		"200URL": {
			url:      fmt.Sprintf("https://httpstat.us/%v", http.StatusOK),
			response: http.StatusOK,
		},
		"400URL": {
			url:      fmt.Sprintf("https://httpstat.us/%v", http.StatusBadRequest),
			response: http.StatusBadRequest,
		},
		"500URL": {
			url:      fmt.Sprintf("https://httpstat.us/%v", http.StatusInternalServerError),
			response: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(fmt.Sprintf("TestMetricsWith%s", name), func(t *testing.T) {
			t.Parallel()

			r := http.NewServeMux()

			r.HandleFunc("/", Scraper())
			r.Handle("/metrics", metrics.NewMetricsHandler())

			req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(fmt.Sprintf(`{"url":"%s"}`, test.url))))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusCreated, rr.Code)

			time.Sleep(2 * time.Second)

			req, err = http.NewRequest(http.MethodGet, "/metrics", nil)
			if err != nil {
				t.Fatal(err)
			}

			r.ServeHTTP(rr, req)

			assert.Contains(t, rr.Body.String(), fmt.Sprintf(`http_get{code="%v",url="%s"}`, test.response, test.url))
		})
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(health.HealthGet())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Contains(t, rr.Body.String(), "status")
	assert.Contains(t, rr.Body.String(), "timestamp")
	assert.Contains(t, rr.Body.String(), "component")
	assert.Contains(t, rr.Body.String(), "name")
	assert.Contains(t, rr.Body.String(), "version")
}
