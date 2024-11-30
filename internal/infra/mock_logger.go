package infra

import (
	"github.com/stretchr/testify/mock"
)

// MockLogger is a mock of the Logger interface for testing purposes
type MockLogger struct {
	mock.Mock
}

// Info mocks the Info method
func (m *MockLogger) Info(msg string) {
	m.Called(msg)
}

// Error mocks the Error method
func (m *MockLogger) Error(msg string) {
	m.Called(msg)
}
