package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthGet())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Contains(t, rr.Body.String(), "status")
	assert.Contains(t, rr.Body.String(), "timestamp")
	assert.Contains(t, rr.Body.String(), "component")
	assert.Contains(t, rr.Body.String(), "name")
	assert.Contains(t, rr.Body.String(), "version")
}
