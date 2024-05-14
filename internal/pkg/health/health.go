package health

import (
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

func HealthGet() http.HandlerFunc {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    "phaidra-assessment",
		Version: "0.1",
	}))
	return h.HandlerFunc
}
