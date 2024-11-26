package handler

import (
	"net/http"
	"receipt-processor/internal/logger"

	"github.com/sirupsen/logrus"
)

// HealthCheck handles GET requests on the /health endpoint.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Log the health check access
	logger.Info("Health Check accessed", logrus.Fields{
		"endpoint": "/health",
		"method":   r.Method,
	})

	// Respond with status OK
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		logger.Error("Failed to write Health Check response", logrus.Fields{
			"endpoint": "/health",
			"error":    err,
		})
	}
}
