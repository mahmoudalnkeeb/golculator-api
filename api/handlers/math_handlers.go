package math_handlers

import (
	api_utils "math_api/api/utils/api"
	math_utils "math_api/api/utils/math"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := math_utils.ParseQueryParams(r)
	if err == nil {
		api_utils.SendResultJson(w, math_utils.Add(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := math_utils.ParseQueryParams(r)
	if err == nil {
		api_utils.SendResultJson(w, math_utils.Subtract(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := math_utils.ParseQueryParams(r)
	if err == nil {
		api_utils.SendResultJson(w, math_utils.Multiply(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := math_utils.ParseQueryParams(r)
	if err == nil {
		if n2 == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		api_utils.SendResultJson(w, math_utils.Divide(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}
