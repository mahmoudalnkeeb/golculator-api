package golculator_tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mahmoudalnkeeb/golculator-api/api/handlers/math"
)

func checkStatusCode(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int, testName string) {
	if rr.Code != expectedStatus {
		t.Errorf("%s: expected status %v, got %v", testName, expectedStatus, rr.Code)
	}
}

func checkContentType(t *testing.T, rr *httptest.ResponseRecorder, expectedType string, testName string) {
	contentType := rr.Header().Get("Content-Type")
	if contentType != expectedType {
		t.Errorf("%s: expected Content-Type %v, got %v", testName, expectedType, contentType)
	}
}

func checkJSONResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedResult float64, testName string) {
	expectedResponse := map[string]float64{"result": expectedResult}
	var actualResponse map[string]float64
	if err := json.Unmarshal(rr.Body.Bytes(), &actualResponse); err != nil {
		t.Fatalf("%s: error unmarshaling response: %v", testName, err)
	}

	if actualResponse["result"] != expectedResponse["result"] {
		t.Errorf("%s: expected result %v, got %v", testName, expectedResponse["result"], actualResponse["result"])
	}
}

func checkErrorResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedMessage string, testName string) {
	actualMessage := strings.TrimSpace(rr.Body.String())
	if actualMessage != expectedMessage {
		t.Errorf("%s: expected message %v, got %v", testName, expectedMessage, actualMessage)
	}
}

func checkResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int, expectedResult float64, testName string) {
	checkStatusCode(t, rr, expectedStatus, testName)

	if expectedStatus == http.StatusOK {
		checkContentType(t, rr, "application/json", testName)
		checkJSONResponse(t, rr, expectedResult, testName)
	} else {
		checkContentType(t, rr, "text/plain; charset=utf-8", testName)

		// Determine expected message for bad requests
		var expectedMessage string
		if testName == "DivideByZero" {
			expectedMessage = "Division by zero"
		} else {
			expectedMessage = "Invalid query parameters"
		}

		checkErrorResponse(t, rr, expectedMessage, testName)
	}
}

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		url            string
		expectedStatus int
		expectedResult float64
	}{
		{"AddHandler", math.AddHandler, "/add?n1=10&n2=5", http.StatusOK, 15},                // 10 + 5 = 15
		{"SubtractHandler", math.SubtractHandler, "/subtract?n1=10&n2=5", http.StatusOK, 5},  // 10 - 5 = 5
		{"MultiplyHandler", math.MultiplyHandler, "/multiply?n1=10&n2=5", http.StatusOK, 50}, // 10 * 5 = 50
		{"DivideHandler", math.DivideHandler, "/divide?n1=10&n2=5", http.StatusOK, 2},        // 10 / 5 = 2
		{"DivideByZero", math.DivideHandler, "/divide?n1=10&n2=0", http.StatusBadRequest, 0}, // Division by zero
		{"InvalidParams", math.AddHandler, "/add?n1=abc&n2=5", http.StatusBadRequest, 0},     // Invalid params
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new request
			req, _ := http.NewRequest("GET", tt.url, nil)
			rr := httptest.NewRecorder()

			// Call the handler
			tt.handler(rr, req)

			// Check the response
			checkResponse(t, rr, tt.expectedStatus, tt.expectedResult, tt.name)
		})
	}
}
