package utils_test

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mahmoudalnkeeb/golculator-api/api/utils"
)

func TestAdd(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := utils.Add(num1, num2)
	want := 9.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestSubtract(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := utils.Subtract(num1, num2)
	want := 1.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestMultiply(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := utils.Multiply(num1, num2)
	want := 20.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestDivide(t *testing.T) {
	var num1, num2 = 20.00, 4.00
	got := utils.Divide(num1, num2)
	want := 5.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape utils.Shape
		want  float64
	}{
		{utils.Rectangle{4, 2}, (4 + 2) * 2},
		{utils.Circle{Radius: 2}, 2 * math.Pi * 2},
		{utils.Triangle{Side1: 6, Side2: 8, Base: 12, Height: 7.11}, 26.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %0.2f but wanted %0.2f", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape utils.Shape
		want  float64
	}{
		{utils.Rectangle{Width: 4, Height: 2}, 8.00},
		{utils.Circle{2}, math.Pi * 2 * 2},
		{utils.Triangle{6, 8, 12, 7.11}, 0.5 * (12 * 7.11)},
	}

	epsilon := 0.0001

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if math.Abs(got-tt.want) > epsilon {
			t.Errorf("got %0.2f but wanted %0.2f", got, tt.want)
		}
	}
}

func TestSendResultJson(t *testing.T) {
	// Step 1: Create a response recorder (to capture the response)
	rr := httptest.NewRecorder()

	// Step 2: Call the function with a test result
	testResult := 42.0
	utils.SendResultJson(rr, testResult)

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
