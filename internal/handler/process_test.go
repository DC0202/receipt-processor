package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/internal/model"
	"testing"
)

// Mock Dependencies
type MockUtility struct{}

func (m *MockUtility) ReadBody(r *http.Request) ([]byte, error) {
	return []byte(`{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
			{"shortDescription": "Gatorade", "price": "2.25"},
			{"shortDescription": "Gatorade", "price": "2.25"},
			{"shortDescription": "Gatorade", "price": "2.25"},
			{"shortDescription": "Gatorade", "price": "2.25"}
		],
		"total": "9.00"
	}`), nil
}

func (m *MockUtility) ParseJSON(body []byte, target interface{}) error {
	return json.Unmarshal(body, target)
}

// Mock Services
type MockServices struct{}

func (m *MockServices) GenerateHash(receipt model.Receipt) string {
	return "receipt-hash"
}

func (m *MockServices) CheckReceipt(hash string) (string, bool) {
	return "", false
}

func (m *MockServices) CalculatePoints(receipt model.Receipt) (int, string) {
	return 109, "Points breakdown explanation"
}

func (m *MockServices) StoreReceipt(id, hash string, points int, explanation string) {}

func (m *MockServices) GenerateID() string {
	return "unique-id"
}

// Test ProcessReceipt
func TestProcessReceipt(t *testing.T) {
	mockUtility := &MockUtility{}
	mockServices := &MockServices{}

	req, err := http.NewRequest("POST", "/receipts/process", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	ProcessReceiptWithMocks(rr, req, mockUtility, mockServices)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ProcessReceipt handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":"unique-id"}`
	if rr.Body.String() != expected {
		t.Errorf("ProcessReceipt handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Handler with Mock Dependencies
func ProcessReceiptWithMocks(
	w http.ResponseWriter,
	r *http.Request,
	utility Utility,
	services Services,
) {
	body, err := utility.ReadBody(r)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var receipt model.Receipt
	if err := utility.ParseJSON(body, &receipt); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	receiptHash := services.GenerateHash(receipt)
	if id, exists := services.CheckReceipt(receiptHash); exists {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"` + id + `"}`))
		return
	}

	id := services.GenerateID()
	points, explanation := services.CalculatePoints(receipt)
	services.StoreReceipt(id, receiptHash, points, explanation)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id":"` + id + `"}`))
}
