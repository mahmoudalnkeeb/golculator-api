package api_utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendResultJson(t *testing.T) {
	// Step 1: Create a response recorder (to capture the response)
	rr := httptest.NewRecorder()

	// Step 2: Call the function with a test result
	testResult := 42.0
	SendResultJson(rr, testResult)

	// Step 3: Check the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", rr.Code)
	}

	// Step 4: Check the content type is application/json
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %v", contentType)
	}

	// Step 5: Check the response body contains the expected JSON
	expectedResponse := map[string]float64{"result": testResult}
	var actualResponse map[string]float64
	if err := json.Unmarshal(rr.Body.Bytes(), &actualResponse); err != nil {
		t.Fatalf("error unmarshaling response: %v", err)
	}

	if actualResponse["result"] != expectedResponse["result"] {
		t.Errorf("expected result %v, got %v", expectedResponse["result"], actualResponse["result"])
	}
}
