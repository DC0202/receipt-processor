package main

import (
	"log"
	"os"
	"receipt-processor/internal/logger"
	"receipt-processor/internal/server"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	logger.InitLogger(logLevel)

	if err := server.Run(); err != nil {
		logger.Error("Failed to start server", logrus.Fields{
			"error": err,
		})
	} else {
		logger.Info("Server started successfully", logrus.Fields{})
	}
}
