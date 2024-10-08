package math

import (
	"net/http"

	"github.com/mahmoudalnkeeb/golculator-api/api/utils"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := utils.ParseQueryParams(r)
	if err == nil {
		utils.SendResultJson(w, utils.Subtract(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}
