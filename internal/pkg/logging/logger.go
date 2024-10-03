package logging

import (
	"github.com/spf13/viper"
	"log/slog"
	"os"
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
func InitLogger() {
	once.Do(func() {
		slogLogger, err := NewLogger()
		if err != nil {
			slog.Error("Failed to initialize logger", slog.String("error", err.Error()))
			return
		}
		Log = slogLogger
	})
}

func NewLogger() (*Logger, error) {
	logFilePath := viper.GetString("log_path")

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
			panic(err)
		}
	}
}
