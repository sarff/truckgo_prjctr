package logging

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestInitLoggerError(t *testing.T) {
	tempDir := t.TempDir()
	logFilePath := filepath.Join(tempDir, "testlog.log")
	viper.Set("log_path", logFilePath)
	viper.Set("log_level", "info")

	logger, err := InitLogger(tempDir)
	if err == nil {
		t.Fatal("expected an error due to invalid path, but got none")
	}
	if logger != nil {
		t.Fatalf("expected logger to be nil on error, but got: %v", logger)
	}
}
