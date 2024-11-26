package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"receipt-processor/internal/config"
	"receipt-processor/internal/handler"
	"receipt-processor/internal/logger"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Run sets up and starts the HTTP server.
func Run() error {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize router and handlers
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/receipts/process", handler.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handler.GetPoints).Methods("GET")
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// Configure HTTP server
	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	// Channel to handle system signals for graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("Server starting", logrus.Fields{
			"port": cfg.AppPort,
		})
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error starting server", logrus.Fields{
				"error": err,
			})
		}
	}()

	<-stopChan
	logger.Info("Shutting down server", logrus.Fields{})

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", logrus.Fields{
			"error": err,
		})
		return err
	}

	logger.Info("Server gracefully stopped", logrus.Fields{})
	return nil
}

// loggingMiddleware logs the HTTP requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("HTTP request received", logrus.Fields{
			"method": r.Method,
			"uri":    r.RequestURI,
		})
		next.ServeHTTP(w, r)
	})
}
