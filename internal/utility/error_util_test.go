package utility

import (
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	WriteJSON(responseRecorder, map[string]string{"status": "success"})

	expectedContentType := "application/json"
	if contentType := responseRecorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("WriteJSON() Content-Type = %v; want %v", contentType, expectedContentType)
	}

	// Adjusted to include the newline added by json.Encoder.Encode
	expectedBody := `{"status":"success"}` + "\n"
	if body := responseRecorder.Body.String(); body != expectedBody {
		t.Errorf("WriteJSON() Body = %q; want %q", body, expectedBody)
	}
}

func TestParseJSON(t *testing.T) {
	input := `{"name":"Test"}`
	expected := map[string]string{"name": "Test"}
	var result map[string]string

	err := ParseJSON([]byte(input), &result)
	if err != nil {
		t.Errorf("ParseJSON() error = %v", err)
	}

	if result["name"] != expected["name"] {
		t.Errorf("ParseJSON() got = %v, want %v", result, expected)
	}
}
