package logging

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestInitLogger(t *testing.T) {
	tempDir := t.TempDir()
	logFilePath := filepath.Join(tempDir, "testlog.log")
	viper.Set("log_path", logFilePath)
	viper.Set("log_level", "info")

	logger, err := InitLogger("test")
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
	if logger == nil {
		t.Fatal("expected logger to be initialized, but got nil")
	}

	logFilePath = filepath.Join(tempDir, "test_testlog.log")
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		t.Fatalf("expected log file to be created, but it was not: %v", err)
	}

	err = os.Remove(logFilePath)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}

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
