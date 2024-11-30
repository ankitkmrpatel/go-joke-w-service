package infra

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// Logger interface that abstracts logrus methods.
type Logger interface {
	Info(msg string)
	Error(msg string)
	// Add other methods like Warn, Debug, etc.
}

type LoggerImpl struct {
	log    *logrus.Logger
	file   *os.File
	logMux sync.Mutex
}

// NewLogger initializes a new Logger.
func NewLogger(filePath string) *LoggerImpl {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &LoggerImpl{
		log:  logger,
		file: file,
	}
}

// Info logs an informational message.
func (l *LoggerImpl) Info(message string) {
	l.logMux.Lock()
	defer l.logMux.Unlock()
	l.log.Info(message)
}

// Info logs an informational message.
func (l *LoggerImpl) Error(message string) {
	l.logMux.Lock()
	defer l.logMux.Unlock()
	l.log.Error(message)
}

// UpdateLogFile updates the logger to use a new log file.
func (l *LoggerImpl) UpdateLogFile(newFilePath string) {
	l.logMux.Lock()
	defer l.logMux.Unlock()

	newFile, err := os.OpenFile(newFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		l.log.Errorf("Failed to update log file: %v", err)
		return
	}

	// Close the old file
	l.file.Close()
	l.file = newFile
	l.log.SetOutput(newFile)
}

// Close closes the logger's file handle.
func (l *LoggerImpl) Close() {
	l.logMux.Lock()
	defer l.logMux.Unlock()
	l.file.Close()
}
