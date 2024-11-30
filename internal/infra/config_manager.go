package infra

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/models"

	"github.com/fsnotify/fsnotify"
)

// LoadConfig reads the configuration from the specified path and returns it as a pointer.
func LoadConfig(path string) (*models.Config, error) {
	var cfg models.Config
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// WatchConfig monitors changes to the config file and triggers the provided callback on updates.
func WatchConfig(path string, onUpdate func(*models.Config)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to initialize config watcher: %v", err)
	}
	defer watcher.Close()

	err = watcher.Add(path)
	if err != nil {
		log.Fatalf("Failed to watch config file: %v", err)
	}

	// Use a debounce mechanism to prevent frequent reloads
	var mu sync.Mutex
	var lastUpdate time.Time
	const debounceDuration = 500 * time.Millisecond

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				mu.Lock()
				if time.Since(lastUpdate) > debounceDuration {
					lastUpdate = time.Now()
					mu.Unlock()

					newCfg, err := LoadConfig(path)
					if err == nil {
						onUpdate(newCfg) // Trigger callback with the updated config
					} else {
						log.Printf("Failed to load updated config: %v", err)
					}
				} else {
					mu.Unlock()
				}
			}
		case err := <-watcher.Errors:
			log.Printf("Config watcher error: %v", err)
		}
	}
}
