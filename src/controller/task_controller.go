package controller

import (
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
	"strconv"

	"github.com/gorilla/mux"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var result string

	method := r.Method
	w.Header().Set("Content-Type", "application/json")

	task_repo := model.NewTaskResitory()

	switch method {
	case "GET":

		vars := mux.Vars(r)
		string_id := vars["id"]
		id, err := strconv.Atoi(string_id)

		if err != nil {
			result = task_repo.GetAll()
		} else {
			result = task_repo.GetOne(id)
		}

		fmt.Fprint(w, result)
		break
	}

}
