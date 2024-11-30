package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/business"
	"github.com/ankitkmrpatel/go-joke-w-service/internal/infra"
	"github.com/ankitkmrpatel/go-joke-w-service/internal/models"
	"github.com/ankitkmrpatel/go-joke-w-service/utils"
)

func main() {
	// Load configuration
	configPath := "config/config.json"
	cfg, err := infra.LoadConfig(configPath)
	utils.HandleError(err, "loading configuration", true) // Call HandleError for critical errors

	// Validate the config
	err = cfg.Validate()
	utils.HandleError(err, "validating configuration", true) // Call HandleError for critical errors

	// Initialize logger
	logger := infra.NewLogger(cfg.LogFilePath)
	defer logger.Close()

	// Initialize metrics server
	infra.InitMetricsServer(cfg.MetricsServer)

	// Start configuration watcher
	go infra.WatchConfig(configPath, func(newCfg *models.Config) {
		cfg.Lock()
		defer cfg.Unlock()

		if newCfg.LogFilePath != cfg.LogFilePath {
			logger.Info("Log file path changed. Updating logger.")
			logger.UpdateLogFile(newCfg.LogFilePath)
		}

		// Safely update the config fields
		cfg.Jokes = newCfg.Jokes
		cfg.MetricsServer = newCfg.MetricsServer

		infra.ConfigReloadCounter.Inc()
		logger.Info("Configuration updated successfully.")
	})

	// Initialize business services
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Start 10-second timer for fetching jokes from config
	go func() {
		defer wg.Done() // Ensure Done is called once
		business.StartTimer(10*time.Second, logger, func() {
			business.PrintRandomJoke(cfg, logger)
			infra.JokesFromConfigCounter.Inc()
		})
	}()

	// Start 15-second timer for fetching jokes from API
	go func() {
		defer wg.Done() // Ensure Done is called once
		business.StartTimer(15*time.Second, logger, func() {
			business.PrintAPIJoke(logger)
			infra.JokesFromAPICounter.Inc()
		})
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down gracefully...")
	wg.Wait()
}
