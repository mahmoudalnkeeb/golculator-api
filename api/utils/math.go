package utils

import (
	"errors"
	"math"
	"net/http"
	"strconv"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Side1  float64
	Side2  float64
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Base
}

func Add(f1 float64, f2 float64) float64 {
	return f1 + f2
}

func Subtract(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func ParseQueryParams(r *http.Request) (float64, float64, error) {
	n1String := r.URL.Query().Get("n1")
	n2String := r.URL.Query().Get("n2")

	n1, err1 := strconv.ParseFloat(n1String, 64)
	n2, err2 := strconv.ParseFloat(n2String, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid parameters")
	}

	return n1, n2, nil
}
