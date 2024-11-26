package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type MockPointsServices struct{}

func (m *MockPointsServices) GetReceiptPoints(id string, detailed bool) (int, string, bool) {
	if id == "id12345" {
		return 109, "Points breakdown explanation", true
	}
	return 0, "", false
}

func TestGetPoints(t *testing.T) {
	mockServices := &MockPointsServices{}

	// Test non-detailed response
	req, err := http.NewRequest("GET", "/receipts/id12345/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "id12345"})

	rr := httptest.NewRecorder()
	GetPointsWithMocks(rr, req, mockServices)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPoints handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := map[string]interface{}{
		"points": 109,
	}
	expectedBody, _ := json.Marshal(expectedResponse)
	if rr.Body.String() != string(expectedBody)+"\n" {
		t.Errorf("GetPoints handler returned unexpected body: got %v want %v", rr.Body.String(), string(expectedBody)+"\n")
	}

	// Test detailed response
	req, err = http.NewRequest("GET", "/receipts/id12345/points?detailed=true", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "id12345"})

	rr = httptest.NewRecorder()
	GetPointsWithMocks(rr, req, mockServices)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetPoints handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedDetailedResponse := map[string]interface{}{
		"points":      109,
		"explanation": "Points breakdown explanation",
	}
	expectedDetailedBody, _ := json.Marshal(expectedDetailedResponse)
	if rr.Body.String() != string(expectedDetailedBody)+"\n" {
		t.Errorf("GetPoints handler returned unexpected body: got %v want %v", rr.Body.String(), string(expectedDetailedBody)+"\n")
	}
}

// Handler with Mock Dependencies
func GetPointsWithMocks(
	w http.ResponseWriter,
	r *http.Request,
	services PointsServices,
) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Missing receipt ID", http.StatusBadRequest)
		return
	}

	detailed := r.URL.Query().Get("detailed") == "true"
	points, explanation, found := services.GetReceiptPoints(id, detailed)
	if !found {
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{"points": points}
	if detailed {
		response["explanation"] = explanation
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
