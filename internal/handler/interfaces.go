package handler

import (
	"net/http"
	"receipt-processor/internal/model"
)

// Utility interface for utility functions
type Utility interface {
	ReadBody(r *http.Request) ([]byte, error)
	ParseJSON(body []byte, target interface{}) error
}

// Services interface for receipt services
type Services interface {
	GenerateHash(receipt model.Receipt) string
	CheckReceipt(hash string) (string, bool)
	CalculatePoints(receipt model.Receipt) (int, string)
	StoreReceipt(id, hash string, points int, explanation string)
	GenerateID() string
}

// PointsServices interface for points retrieval
type PointsServices interface {
	GetReceiptPoints(id string, detailed bool) (int, string, bool)
}
