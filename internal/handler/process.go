package handler

import (
	"net/http"
	"receipt-processor/internal/logger"
	"receipt-processor/internal/model"
	"receipt-processor/internal/services"
	"receipt-processor/internal/utility"

	"github.com/sirupsen/logrus"
)

// ProcessReceipt handles POST requests on the /process endpoint for receipt processing.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := utility.ReadBody(r)
	if err != nil {
		logger.Error("Failed to read request body", logrus.Fields{
			"error":    err,
			"endpoint": "/process",
		})
		utility.WriteError(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var dataMap map[string]interface{}
	if err := utility.ParseJSON(body, &dataMap); err != nil {
		logger.Error("Invalid JSON format in request", logrus.Fields{
			"error":    err,
			"endpoint": "/process",
		})
		utility.WriteError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	var receipt model.Receipt
	if err := receipt.ValidateReceiptMap(dataMap); err != nil {
		logger.Error("Invalid receipt data in request", logrus.Fields{
			"error":    err,
			"endpoint": "/process",
		})
		utility.WriteError(w, "Incorrect Receipt data", http.StatusBadRequest)
		return
	}

	if err := utility.ParseJSON(body, &receipt); err != nil {
		logger.Error("Failed to parse receipt JSON", logrus.Fields{
			"error":    err,
			"endpoint": "/process",
		})
		utility.WriteError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if err := receipt.Validate(); err != nil {
		logger.Error("Receipt validation failed", logrus.Fields{
			"error":    err,
			"endpoint": "/process",
			"receipt":  receipt,
		})
		utility.WriteError(w, "Validation error", http.StatusBadRequest)
		return
	}

	// Process receipt hash and check existence
	receiptHash := services.GenerateHash(receipt)
	if id, exists := services.CheckReceipt(receiptHash); exists {
		logger.Info("Receipt already processed", logrus.Fields{
			"id":       id,
			"endpoint": "/process",
		})
		utility.WriteJSON(w, map[string]string{"id": id})
		return
	}

	// Generate ID, calculate points, and store receipt
	id := utility.GenerateID()
	points, explanation := services.CalculatePoints(receipt)
	services.StoreReceipt(id, receiptHash, points, explanation)

	logger.Info("Receipt processed successfully", logrus.Fields{
		"id":       id,
		"points":   points,
		"endpoint": "/process",
	})
	utility.WriteJSON(w, map[string]string{"id": id})
}
