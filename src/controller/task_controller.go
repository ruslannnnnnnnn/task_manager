package controller

import (
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
	"strconv"

	"github.com/gorilla/mux"
)

func TaskGetHandler(w http.ResponseWriter, r *http.Request) {
	result := "{}"

	w.Header().Set("Content-Type", "application/json")

	task_repo := model.NewTaskResitory()

	vars := mux.Vars(r)
	string_id := vars["id"]
	id, err := strconv.Atoi(string_id)

	if string_id == "" {
		result = task_repo.GetAll()
	} else if err == nil {
		result = task_repo.GetOne(id)
	}

	fmt.Fprint(w, result)
}

func TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task_repo := model.NewTaskResitory()

	fmt.Fprint(w, task_repo.Post(r.Body))
}
