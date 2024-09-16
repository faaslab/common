package logger

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {

	l.Infow("Failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	l.Infof("Failed to fetch URL: %s", "http://example.com")

	if l == nil {
		t.Error("Expected non-nil logger, got nil")
	}

	// Check if the returned logger is of type *zap.SugaredLogger
	_, ok := interface{}(l).(*zap.SugaredLogger)
	if !ok {
		t.Error("Expected *zap.SugaredLogger, got different type")
	}
}
