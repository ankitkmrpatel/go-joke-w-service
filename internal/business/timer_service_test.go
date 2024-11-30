package business

import (
	"testing"
	"time"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/infra"
	"github.com/stretchr/testify/mock"
)

func TestStartTimer(t *testing.T) {
	mockLogger := new(infra.MockLogger)
	// cfg := &models.Config{Jokes: []string{"Joke 1", "Joke 2"}}

	// Mock the Info method to avoid actual logging during tests
	mockLogger.On("Info", mock.Anything).Return()

	// Start a timer that will run the task once
	go StartTimer(1*time.Second, mockLogger, func() {
		mockLogger.Info("Timer triggered")
	})

	// Wait to give the timer a chance to trigger
	time.Sleep(2 * time.Second)

	// Verify that the timer ran and called the mock logger
	mockLogger.AssertExpectations(t)
}
