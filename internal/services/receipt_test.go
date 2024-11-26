package services

import (
	"receipt-processor/internal/model"
	"testing"
)

// Test GenerateHash
func TestGenerateHash(t *testing.T) {
	receipt := model.Receipt{
		Retailer: "M&M Corner Market",
		Total:    "9.00",
	}
	hash := GenerateHash(receipt)

	if hash == "" {
		t.Errorf("expected a non-empty hash for the receipt")
	}

	// Ensure different receipts generate different hashes
	anotherReceipt := model.Receipt{
		Retailer: "Another Market",
		Total:    "10.00",
	}
	anotherHash := GenerateHash(anotherReceipt)

	if hash == anotherHash {
		t.Errorf("expected different hashes for different receipts")
	}
}

// Test CheckReceipt
func TestCheckReceipt(t *testing.T) {
	receiptHash := "sampleHash123"
	receipts[receiptHash] = "12345"

	id, found := CheckReceipt(receiptHash)
	if !found {
		t.Errorf("expected receipt to be found")
	}
	if id != "12345" {
		t.Errorf("expected receipt ID to be '12345'; got %s", id)
	}

	// Test for a hash that doesn't exist
	id, found = CheckReceipt("nonexistentHash")
	if found {
		t.Errorf("expected receipt to not be found")
	}
}

// Test StoreReceipt
func TestStoreReceipt(t *testing.T) {
	id := "12345"
	hash := "sampleHash123"
	points := 109
	explanation := "Points breakdown explanation"

	StoreReceipt(id, hash, points, explanation)

	// Verify storage
	if receipts[hash] != id {
		t.Errorf("expected receipt hash to map to ID %s", id)
	}

	details, ok := receiptDetails[id]
	if !ok {
		t.Errorf("expected receipt details to be stored for ID %s", id)
	}

	if details.Points != points {
		t.Errorf("expected points to be %d; got %d", points, details.Points)
	}
	if details.Explanation != explanation {
		t.Errorf("expected explanation to be %q; got %q", explanation, details.Explanation)
	}
}

// Test GetReceiptPoints
func TestGetReceiptPoints(t *testing.T) {
	id := "12345"
	points := 109
	explanation := "Points breakdown explanation"
	receiptDetails[id] = model.ReceiptDetails{
		Points:      points,
		Explanation: explanation,
	}

	// Test retrieving points with detailed explanation
	retrievedPoints, retrievedExplanation, found := GetReceiptPoints(id, true)
	if !found {
		t.Errorf("expected receipt to be found")
	}
	if retrievedPoints != points {
		t.Errorf("expected points to be %d; got %d", points, retrievedPoints)
	}
	if retrievedExplanation != explanation {
		t.Errorf("expected explanation to be %q; got %q", explanation, retrievedExplanation)
	}

	// Test retrieving points without detailed explanation
	retrievedPoints, retrievedExplanation, found = GetReceiptPoints(id, false)
	if retrievedExplanation != "" {
		t.Errorf("expected no explanation; got %q", retrievedExplanation)
	}

	// Test for nonexistent receipt ID
	_, _, found = GetReceiptPoints("nonexistentID", true)
	if found {
		t.Errorf("expected receipt to not be found")
	}
}
