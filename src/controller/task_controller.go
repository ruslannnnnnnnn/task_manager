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
	statusCode := 200

	w.Header().Set("Content-Type", "application/json")

	task_repo := model.NewTaskRepository()

	vars := mux.Vars(r)
	string_id := vars["id"]
	id, err := strconv.Atoi(string_id)

	if string_id == "" {
		result, statusCode = task_repo.GetAll()
	} else if err == nil {
		result, statusCode = task_repo.GetOne(id)
	} else {
		result = `{"error":"idk bruh"}`
		statusCode = 500
	}

	w.WriteHeader(statusCode)
	fmt.Fprint(w, result)
}

func TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task_repo := model.NewTaskRepository()

	var result string
	var statusCode int
	result, statusCode = task_repo.Post(r.Body)

	w.WriteHeader(statusCode)
	fmt.Fprint(w, result)
}
