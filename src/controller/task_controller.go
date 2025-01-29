package controller

import (
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
	"strconv"

	"github.com/gorilla/mux"
)

func TaskGetHandler(w http.ResponseWriter, r *http.Request) {
	var result *model.ApiResponse
	w.Header().Set("Content-Type", "application/json")

	task_repo := model.NewTaskRepository()

	vars := mux.Vars(r)
	string_id := vars["id"]
	id, err := strconv.Atoi(string_id)

	if string_id == "" {
		result = task_repo.GetAll(500, 0)
	} else if err == nil {
		result = task_repo.GetOne(id)
	}

	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}

func TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task_repo := model.NewTaskRepository()

	result := task_repo.Post(r.Body)

	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}

func TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task_repo := model.NewTaskRepository()

	vars := mux.Vars(r)
	string_id := vars["id"]
	id, err := strconv.Atoi(string_id)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprint(w, "{\"error\":\"id is invalid.\"}")
		return
	}

	result := task_repo.Delete(id)
	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}
