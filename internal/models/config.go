package models

import (
	"sync"
)

// Config represents the structure of the configuration file.
type Config struct {
	Jokes         []string `json:"jokes"`
	LogFilePath   string   `json:"log_file_path"`
	MetricsServer string   `json:"metrics_server"`
	mu            sync.RWMutex
}

// Lock the config for writing.
func (c *Config) Lock() {
	c.mu.Lock()
}

// Unlock the config after writing.
func (c *Config) Unlock() {
	c.mu.Unlock()
}

// RLock locks the config for reading.
func (c *Config) RLock() {
	c.mu.RLock()
}

// RUnlock unlocks the config after reading.
func (c *Config) RUnlock() {
	c.mu.RUnlock()
}
