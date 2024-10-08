package api_utils

import (
	"encoding/json"
	"net/http"
)

func SendResultJson(w http.ResponseWriter, result float64) {

	response := map[string]float64{"result": result}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
