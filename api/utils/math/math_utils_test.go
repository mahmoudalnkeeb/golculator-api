package math_utils

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := Add(num1, num2)
	want := 9.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestSubtract(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := Subtract(num1, num2)
	want := 1.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestMultiply(t *testing.T) {
	var num1, num2 = 5.00, 4.00
	got := Multiply(num1, num2)
	want := 20.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestDivide(t *testing.T) {
	var num1, num2 = 20.00, 4.00
	got := Divide(num1, num2)
	want := 5.00

	if got != want {
		t.Errorf("got %0.2f but wanted %0.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{4, 2}, (4 + 2) * 2},
		{Circle{2}, 2 * math.Pi * 2},
		{Triangle{6, 8, 12, 7.11}, 26.0},
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
		shape Shape
		want  float64
	}{
		{Rectangle{4, 2}, 8.00},
		{Circle{2}, math.Pi * 2 * 2},
		{Triangle{6, 8, 12, 7.11}, 0.5 * (12 * 7.11)},
	}

	epsilon := 0.0001

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if math.Abs(got-tt.want) > epsilon {
			t.Errorf("got %0.2f but wanted %0.2f", got, tt.want)
		}
	}
}
