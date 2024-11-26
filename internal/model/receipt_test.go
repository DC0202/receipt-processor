package model

import (
	"testing"
)

func TestReceipt_ValidateReceiptMap(t *testing.T) {
	receipt := Receipt{}
	validMap := map[string]interface{}{
		"retailer":     "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items":        []interface{}{},
		"total":        "9.00",
	}
	invalidMap := map[string]interface{}{
		"invalidKey": "value",
	}

	// Test valid map
	err := receipt.ValidateReceiptMap(validMap)
	if err != nil {
		t.Errorf("expected no error for valid map; got %v", err)
	}

	// Test invalid map
	err = receipt.ValidateReceiptMap(invalidMap)
	if err == nil || err.Error() != "incorrect receipt data" {
		t.Errorf("expected error for invalid map; got %v", err)
	}
}

func TestReceipt_Validate(t *testing.T) {
	validReceipt := Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	invalidReceipt := Receipt{
		Retailer:     "",
		PurchaseDate: "invalid-date",
		PurchaseTime: "invalid-time",
		Items: []Item{
			{ShortDescription: "", Price: "invalid-price"},
		},
		Total: "invalid-total",
	}

	// Test valid receipt
	err := validReceipt.Validate()
	if err != nil {
		t.Errorf("expected no error for valid receipt; got %v", err)
	}

	// Test invalid receipt
	err = invalidReceipt.Validate()
	if err == nil {
		t.Errorf("expected error for invalid receipt; got nil")
	}
}

func TestItem_Validate(t *testing.T) {
	validItem := Item{ShortDescription: "Gatorade", Price: "2.25"}
	invalidItem := Item{ShortDescription: "", Price: "invalid-price"}

	// Test valid item
	err := validItem.Validate()
	if err != nil {
		t.Errorf("expected no error for valid item; got %v", err)
	}

	// Test invalid item
	err = invalidItem.Validate()
	if err == nil || err.Error() != "item short description is invalid" {
		t.Errorf("expected error for invalid item; got %v", err)
	}
}

func TestContains(t *testing.T) {
	validSlice := []string{"retailer", "purchaseDate", "items", "total"}
	tests := []struct {
		slice   []string
		element string
		want    bool
	}{
		{validSlice, "retailer", true},
		{validSlice, "unknownKey", false},
	}

	for _, test := range tests {
		got := contains(test.slice, test.element)
		if got != test.want {
			t.Errorf("contains(%v, %q) = %v; want %v", test.slice, test.element, got, test.want)
		}
	}
}
