package infra

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define Prometheus metrics
	JokesFromConfigCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "jokes_from_config_total",
		Help: "Total number of jokes fetched from the configuration file",
	})
	JokesFromAPICounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "jokes_from_api_total",
		Help: "Total number of jokes fetched from the API",
	})
	APIRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "api_request_duration_seconds",
		Help:    "Histogram of API request durations",
		Buckets: prometheus.DefBuckets,
	})
	ConfigReloadCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "config_reload_total",
		Help: "Total number of configuration reloads",
	})
)

// InitMetricsServer starts the metrics server and handles graceful shutdown
func InitMetricsServer(metricsServerURL string) {
	// Set up a context that listens for shutdown signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up the HTTP server
	http.Handle("/metrics", promhttp.Handler())
	server := &http.Server{Addr: metricsServerURL}

	// Start the server in a goroutine
	go func() {
		log.Printf("Metrics server running on %s", metricsServerURL)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Metrics server failed: %v", err)
		}
	}()

	// Set up a channel to listen for OS signals (SIGINT, SIGTERM)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// Wait for a shutdown signal
	<-stop

	// Gracefully shut down the server with a timeout
	log.Println("Shutting down metrics server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Metrics server shutdown failed: %v", err)
	}

	log.Println("Metrics server stopped.")
}
