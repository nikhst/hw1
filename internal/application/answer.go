package application

import (
	"fmt"
	"net/http"

	"github.com/nikhst/rpn/pkg/rpn"
)

func Answer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var expression string = r.FormValue("expression")
	var response float64
	response, err := rpn.Calc(expression)
	if err != nil {
		http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprintf(w, "%g", response)
}
