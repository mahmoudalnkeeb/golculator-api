package main

import (
	"fmt"
	"net/http"

	"github.com/mahmoudalnkeeb/golculator-api/api/handlers/math"
)

func main() {
	http.HandleFunc("/add", math.AddHandler)
	http.HandleFunc("/subtract", math.SubtractHandler)
	http.HandleFunc("/divide", math.DivideHandler)
	http.HandleFunc("/multiply", math.MultiplyHandler)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}
}
