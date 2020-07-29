package log

import (
	"testing"
)

func TestInfof(t *testing.T) {
	Debugf("test log, level: %s", DebugLevel.String())
	Infof("test log, level: %s", InfoLevel.String())
	Warnf("test log, level: %s", WarnLevel.String())
	Errorf("test log, level: %s", ErrorLevel.String())
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger(WithLevel(InfoLevel))
	logger.Logf(DebugLevel, "test log, level: %s", DebugLevel.String())
	logger.Logf(InfoLevel, "test log, level: %s", InfoLevel.String())
	logger.Logf(WarnLevel, "test log, level: %s", WarnLevel.String())
	logger.Logf(ErrorLevel, "test log, level: %s", ErrorLevel.String())
}
