package business

import (
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/infra"
)

// StartTimer initializes and runs a timer at the given interval.
// The `task` parameter is a function that will be executed on each tick.
func StartTimer(interval time.Duration, logger infra.Logger, task func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	logger.Info("Timer started with interval: " + interval.String())

	// Use `for range` to iterate over ticker.C
	for range ticker.C {
		logger.Info("Executing task at interval: " + interval.String())
		task()
	}
}
