package math

import (
	"net/http"

	"github.com/mahmoudalnkeeb/golculator-api/api/utils"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	n1, n2, err := utils.ParseQueryParams(r)
	if err == nil {
		if n2 == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		utils.SendResultJson(w, utils.Divide(n1, n2))
	} else {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
	}
}
