package infra

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestWatchConfig(t *testing.T) {
	// Temporary config file path for testing
	configFile := "/config/test_config.json"

	// Create a test config file with initial data
	initialConfig := models.Config{
		LogFilePath: "/log/temp-file.log",
	}

	file, err := os.Create(configFile)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	defer file.Close()

	// Write initial config to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&initialConfig)
	if err != nil {
		t.Fatalf("Failed to write initial config: %v", err)
	}

	// Function to simulate config update
	onUpdate := func(cfg *models.Config) {
		assert.Equal(t, "/log/temp-file-2.log", cfg.LogFilePath)
	}

	// Run the WatchConfig function in a goroutine so that it doesn't block the test
	go WatchConfig(configFile, onUpdate)

	// Sleep for a brief moment to ensure WatchConfig is ready
	time.Sleep(1 * time.Second)

	// Simulate a change in the config file by modifying the LogFilePath
	newConfig := models.Config{
		LogFilePath: "/log/temp-file-2.log",
	}

	// Open the file and update its content
	file, err = os.OpenFile(configFile, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		t.Fatalf("Failed to open config file for update: %v", err)
	}
	defer file.Close()

	// Write the updated config to the file
	encoder = json.NewEncoder(file)
	err = encoder.Encode(&newConfig)
	if err != nil {
		t.Fatalf("Failed to write updated config: %v", err)
	}

	// Allow time for the watcher to detect the change and trigger the callback
	time.Sleep(1 * time.Second)
}
