// Package config implements functions used to get the default
// configs or create custom configs for the service.
package config

import (
	"os"
	"time"
)

// Struct for configuration parameters
type Config struct {
	ListenPort            string
	MetricsPort           string
	ScraperRequestTimeout time.Duration
	ServerWriteTimeout    time.Duration
	ServerReadTimeout     time.Duration
	ServerIdleTimeout     time.Duration
}

// Return new config object with either default values or with
// values read from environment variables
func NewConfig() *Config {
	serverWriteTimeout, err := getDurationFromEnv("SERVER_WRITE_TIMEOUT", "15s")
	if err != nil {
		panic(err)
	}
	serverReadTimeout, err := getDurationFromEnv("SERVER_READ_TIMEOUT", "15s")
	if err != nil {
		panic(err)
	}
	serverIdleTimeout, err := getDurationFromEnv("SERVER_IDLE_TIMEOUT", "30s")
	if err != nil {
		panic(err)
	}
	scraperRequestTimeout, err := getDurationFromEnv("SCRAPER_REQUEST_TIMEOUT", "30s")
	if err != nil {
		panic(err)
	}

	return &Config{
		ListenPort:            getEnvOrDefault("SERVER_PORT", "8080"),
		MetricsPort:           getEnvOrDefault("METRICS_PORT", "9095"),
		ServerWriteTimeout:    serverWriteTimeout,
		ServerReadTimeout:     serverReadTimeout,
		ServerIdleTimeout:     serverIdleTimeout,
		ScraperRequestTimeout: scraperRequestTimeout,
	}
}

// Gets an environment variable value or returns set default.
func getEnvOrDefault(name, value string) string {
	v := os.Getenv(name)
	if v == "" {
		v = value
	}
	return v
}

// Gets a environment variable value as a time.Duration value
// or returns a set default.
func getDurationFromEnv(key, value string) (time.Duration, error) {
	v := getEnvOrDefault(key, value)
	return time.ParseDuration(v)
}
