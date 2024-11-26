package services

import (
	"receipt-processor/internal/model"
	"strings"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	// Sample Receipt for Testing
	receipt := model.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []model.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	expectedPoints := 109 // Expected points based on breakdown
	points, explanation := CalculatePoints(receipt)

	if points != expectedPoints {
		t.Errorf("expected %d points; got %d points", expectedPoints, points)
	}

	expectedExplanationSubstring := "50 points - total is a round dollar amount"
	if !containsSubstring(explanation, expectedExplanationSubstring) {
		t.Errorf("expected explanation to contain %q; got %q", expectedExplanationSubstring, explanation)
	}
}

func TestCalculatePoints_InvalidData(t *testing.T) {
	// Receipt with invalid data for edge cases
	receipt := model.Receipt{
		Retailer:     "",
		PurchaseDate: "invalid-date",
		PurchaseTime: "invalid-time",
		Items: []model.Item{
			{ShortDescription: "Gatorade", Price: "invalid-price"},
		},
		Total: "invalid-total",
	}

	points, explanation := CalculatePoints(receipt)

	if points != 0 {
		t.Errorf("expected 0 points for invalid data; got %d", points)
	}

	expectedExplanationSubstring := "Breakdown:"
	if !containsSubstring(explanation, expectedExplanationSubstring) {
		t.Errorf("expected explanation to contain %q; got %q", expectedExplanationSubstring, explanation)
	}
}

// Helper function to check if a string contains a substring
func containsSubstring(str, substring string) bool {
	return strings.Contains(str, substring)
}
