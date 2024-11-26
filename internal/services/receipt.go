package services

import (
	"receipt-processor/internal/logger"
	"receipt-processor/internal/model"
	"receipt-processor/pkg/hash"

	"github.com/sirupsen/logrus"
)

var receipts = make(map[string]string)
var receiptDetails = map[string]model.ReceiptDetails{}

// GenerateHash computes a SHA-1 hash for the given receipt.
func GenerateHash(receipt model.Receipt) string {
	data := receipt.String()
	h := hash.GenerateHash(data)
	logger.Info("Generated hash for receipt", logrus.Fields{
		"hash": h,
	})
	return h
}

// CheckReceipt checks if a receipt hash already exists and returns the corresponding ID.
func CheckReceipt(hash string) (string, bool) {
	id, found := receipts[hash]
	if found {
		logger.Info("Receipt already processed", logrus.Fields{
			"receipt_id": id,
			"hash":       hash,
		})
	} else {
		logger.Warn("New receipt hash detected", logrus.Fields{
			"hash": hash,
		})
	}
	return id, found
}

// StoreReceipt stores the receipt hash and associated details.
func StoreReceipt(id, hash string, points int, explanation string) {
	receipts[hash] = id
	receiptDetails[id] = model.ReceiptDetails{
		Points:      points,
		Explanation: explanation,
	}
	logger.Info("Stored receipt details", logrus.Fields{
		"receipt_id": id,
		"hash":       hash,
		"points":     points,
	})
}

// GetReceiptPoints retrieves points and explanation based on receipt ID.
func GetReceiptPoints(id string, detailed bool) (int, string, bool) {
	if details, ok := receiptDetails[id]; ok {
		logger.Info("Retrieved receipt details", logrus.Fields{
			"receipt_id": id,
			"detailed":   detailed,
		})
		if detailed {
			logger.Info("Returning detailed explanation for receipt", logrus.Fields{
				"receipt_id": id,
			})
			return details.Points, details.Explanation, true
		}
		return details.Points, "", true
	}
	logger.Warn("Receipt not found", logrus.Fields{
		"receipt_id": id,
	})
	return 0, "", false
}
