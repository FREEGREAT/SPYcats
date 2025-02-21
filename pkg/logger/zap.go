// pkg/logger/logger.go

package logger

import (
	"context"
	"os"
	middleware "spy-cats/internal/middlewae" // Оновіть на ваш правильний шлях імпорту

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

// New creates a new logger instance
func New(level string) middleware.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetOutput(os.Stdout)

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	log.SetLevel(lvl)

	return &Logger{
		logger: log,
	}
}

func (l *Logger) Named(name string) middleware.Logger {
	return &Logger{
		logger: l.logger.WithField("logger_name", name).Logger,
	}
}

func (l *Logger) With(args ...interface{}) middleware.Logger {
	if len(args)%2 != 0 {
		l.logger.Error("Invalid number of arguments for With() method")
		return l
	}

	fields := make(logrus.Fields)
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fields[key] = args[i+1]
	}

	return &Logger{
		logger: l.logger.WithFields(fields).Logger,
	}
}

func (l *Logger) WithContext(ctx context.Context) middleware.Logger {
	requestID := ctx.Value("RequestID")
	if requestID != nil {
		return l.With("request_id", requestID)
	}
	return l
}

func (l *Logger) Debug(message string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.WithFields(argsToFields(args...)).Debug(message)
	} else {
		l.logger.Debug(message)
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.WithFields(argsToFields(args...)).Info(message)
	} else {
		l.logger.Info(message)
	}
}

func (l *Logger) Warn(message string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.WithFields(argsToFields(args...)).Warn(message)
	} else {
		l.logger.Warn(message)
	}
}

func (l *Logger) Error(message string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.WithFields(argsToFields(args...)).Error(message)
	} else {
		l.logger.Error(message)
	}
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.WithFields(argsToFields(args...)).Fatal(message)
	} else {
		l.logger.Fatal(message)
	}
}

// Helper function to convert args to logrus Fields
func argsToFields(args ...interface{}) logrus.Fields {
	fields := make(logrus.Fields)

	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}

		key, ok := args[i].(string)
		if !ok {
			continue
		}

		fields[key] = args[i+1]
	}

	return fields
}
