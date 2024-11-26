package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger(logLevel string) {
	// Open the log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// Set log output to both file and console
	multiWriter := io.MultiWriter(file, os.Stdout)
	log.Out = multiWriter

	// Use LOG_LEVEL from environment if not provided
	if logLevel == "" {
		logLevel = os.Getenv("LOG_LEVEL") // Default to environment variable
		if logLevel == "" {
			logLevel = "info" // Fallback to INFO if LOG_LEVEL is not set
		}
	}

	// Parse and set log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Warnf("Invalid log level: %s, defaulting to INFO", logLevel)
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Set log format
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

func Warn(message string, fields logrus.Fields) {
	log.WithFields(fields).Warn(message)
}

func Error(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

func Fatal(message string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(message)
}
