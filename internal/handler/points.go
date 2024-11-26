package handler

import (
	"net/http"
	"receipt-processor/internal/logger"
	"receipt-processor/internal/services"
	"receipt-processor/internal/utility"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetPoints handles GET requests on the /{id}/points endpoint to retrieve points for a specific receipt.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		logger.Error("Missing receipt ID in request", logrus.Fields{
			"endpoint": "/{id}/points",
		})
		utility.WriteError(w, "Missing receipt ID in request", http.StatusBadRequest)
		return
	}

	// Check if the detailed flag is set in the query
	detailed := r.URL.Query().Get("detailed") == "true"

	// Retrieve points and explanation for the receipt
	points, explanation, ok := services.GetReceiptPoints(id, detailed)
	if !ok {
		logger.Error("Invalid receipt ID", logrus.Fields{
			"receipt_id": id,
			"endpoint":   "/{id}/points",
		})
		utility.WriteError(w, "Incorrect receipt ID", http.StatusNotFound)
		return
	}

	// Prepare the response
	response := map[string]interface{}{"points": points}
	if detailed {
		response["explanation"] = explanation
	}

	logger.Info("Points retrieved successfully", logrus.Fields{
		"receipt_id": id,
		"points":     points,
		"detailed":   detailed,
		"endpoint":   "/{id}/points",
	})

	utility.WriteJSON(w, response)
}
