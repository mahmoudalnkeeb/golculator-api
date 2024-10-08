package main

import (
	"fmt"
	"net/http"

	math_handlers "math_api/api/handlers"
)

func main() {
	http.HandleFunc("/add", math_handlers.AddHandler)
	http.HandleFunc("/subtract", math_handlers.SubtractHandler)
	http.HandleFunc("/divide", math_handlers.DivideHandler)
	http.HandleFunc("/multiply", math_handlers.MultiplyHandler)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}
}
