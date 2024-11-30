package models

import (
	"errors"
	"fmt"
	"net/url"
)

// ValidateConfig checks if the configuration is valid
func (cfg *Config) Validate() error {
	if len(cfg.Jokes) == 0 {
		return errors.New("jokes list cannot be empty")
	}
	// Validate LogFilePath
	if cfg.LogFilePath == "" {
		return fmt.Errorf("log file path is required")
	}

	// Validate Jokes
	if len(cfg.Jokes) <= 100 {
		return fmt.Errorf("at least 100 jokes are required")
	}

	// Validate MetricsServer URL
	_, err := url.ParseRequestURI(cfg.MetricsServer)
	if err != nil {
		return fmt.Errorf("invalid MetricsServer URL: %v", err)
	}

	return nil
}
