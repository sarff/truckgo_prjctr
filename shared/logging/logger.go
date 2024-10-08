package logging

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// for global use
var (
	Log  LoggerInterface
	once sync.Once
)

type LoggerInterface interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Close()
}

type Logger struct {
	log     *slog.Logger
	logFile *os.File
}

// InitLogger - for global use
func InitLogger(prefix string) (*Logger, error) {
	var initError error
	var slogLogger *Logger
	once.Do(func() {
		slogLogger, initError = newLogger(prefix)
		if initError != nil {
			slog.Error("Failed to initialize logger", slog.String("error", initError.Error()))
		}
	})
	return slogLogger, initError
}

func newLogger(prefix string) (*Logger, error) {
	logFilePath := viper.GetString("log_path")
	dir := filepath.Dir(logFilePath)
	fileName := filepath.Base(logFilePath)
	extension := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, extension)
	newFileName := fmt.Sprintf("%s_%s%s", prefix, nameWithoutExt, extension) // "AUTH_logs.json"
	newPath := filepath.Join(dir, newFileName)

	logFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &Logger{
		log:     slog.New(handler),
		logFile: logFile,
	}, nil
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Close() {
	if l.logFile != nil {
		err := l.logFile.Close()
		if err != nil {
			l.log.Error("Failed to close log file", slog.String("error", err.Error()))
		}
	}
}
