package business

import (
	"testing"

	"github.com/ankitkmrpatel/go-joke-w-service/internal/infra"
	"github.com/ankitkmrpatel/go-joke-w-service/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPrintRandomJoke(t *testing.T) {
	cfg := &models.Config{
		Jokes: []string{"Why do programmers prefer dark mode?", "What is an algorithm?"}, // Test data
	}
	mockLogger := new(infra.MockLogger)

	// Test the function
	mockLogger.On("Info", mock.Anything).Return(nil)
	PrintRandomJoke(cfg, mockLogger)
	mockLogger.AssertExpectations(t)

	// Ensure the logger was called
	assert.True(t, mockLogger.Calls[0].Method == "Info")
}

func TestPrintAPIJoke(t *testing.T) {
	mockLogger := new(infra.MockLogger)

	// Simulate API fetching joke
	mockLogger.On("Info", mock.Anything).Return(nil)
	PrintAPIJoke(mockLogger)
	mockLogger.AssertExpectations(t)

	// Ensure the logger was called
	assert.True(t, mockLogger.Calls[0].Method == "Info")
}
