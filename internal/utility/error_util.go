package utility

import (
	"encoding/json"
	"net/http"
	"receipt-processor/internal/logger"

	"github.com/sirupsen/logrus"
)

// WriteError sends a JSON formatted error message to the client.
func WriteError(w http.ResponseWriter, errMsg string, statusCode int) {
	resp := map[string]string{"error": errMsg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Error("JSON encoding error", logrus.Fields{
			"error_message":  errMsg,
			"status_code":    statusCode,
			"encoding_error": err,
		})
	} else {
		logger.Info("Error response sent", logrus.Fields{
			"error_message": errMsg,
			"status_code":   statusCode,
		})
	}
}

// WriteJSON sends a JSON formatted response to the client.
func WriteJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("JSON encoding error", logrus.Fields{
			"response_data":  data,
			"encoding_error": err,
		})
	} else {
		logger.Info("JSON response sent", logrus.Fields{
			"response_data": data,
		})
	}
}
