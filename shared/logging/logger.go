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

var (
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
	err := os.MkdirAll(dir, os.ModePerm) // os.ModePerm забезпечує повний доступ до директорій
	if err != nil {
		return nil, fmt.Errorf("cant create dir: %w", err)
	}

	logLevel := viper.GetString("log_level")
	var level slog.Level
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	fileName := filepath.Base(logFilePath)
	extension := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, extension)
	newFileName := fmt.Sprintf("%s_%s%s", prefix, nameWithoutExt, extension)
	newPath := filepath.Join(dir, newFileName)

	logFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	var handler slog.Handler
	if level == slog.LevelDebug {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewJSONHandler(logFile, &slog.HandlerOptions{
			Level: level,
		})
	}

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
