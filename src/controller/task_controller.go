package controller

import (
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var result string

	method := r.Method
	w.Header().Set("Content-Type", "application/json")

	switch method {
	case "GET":
		result = model.Get()
		fmt.Fprint(w, result)
	}

}
