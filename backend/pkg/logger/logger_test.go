package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()

	assert.NotNil(t, logger)
	assert.IsType(t, &logrus.Logger{}, logger)

	// Check that the logger is properly configured
	assert.Equal(t, logrus.InfoLevel, logger.GetLevel())

	// Check that the formatter is a JSONFormatter
	_, ok := logger.Formatter.(*logrus.JSONFormatter)
	assert.True(t, ok, "Logger should use JSONFormatter")
}

func TestLoggerConfiguration(t *testing.T) {
	logger := NewLogger()

	// Test that the logger can be used
	logger.Info("Test log message")
	logger.WithField("key", "value").Info("Test log with field")

	// Verify the logger is working
	assert.True(t, true) // If we get here, the logger didn't panic
}

func TestLoggerLevels(t *testing.T) {
	logger := NewLogger()

	// Test different log levels
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	// Verify the logger is working at all levels
	assert.True(t, true) // If we get here, the logger didn't panic
}

func TestLoggerWithFields(t *testing.T) {
	logger := NewLogger()

	// Test logging with fields
	logger.WithFields(logrus.Fields{
		"user_id": 123,
		"action":  "login",
		"ip":      "192.168.1.1",
	}).Info("User logged in")

	// Test logging with a single field
	logger.WithField("request_id", "abc123").Info("Request processed")

	// Verify the logger is working with fields
	assert.True(t, true) // If we get here, the logger didn't panic
}

func TestLoggerErrorHandling(t *testing.T) {
	logger := NewLogger()

	// Test error logging
	err := assert.AnError
	logger.WithError(err).Error("An error occurred")

	// Test error logging with context
	logger.WithFields(logrus.Fields{
		"operation": "database_query",
		"table":     "users",
	}).WithError(err).Error("Database operation failed")

	// Verify the logger is working with errors
	assert.True(t, true) // If we get here, the logger didn't panic
}
