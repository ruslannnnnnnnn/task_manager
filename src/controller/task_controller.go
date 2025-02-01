package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
	"repos/task_manager/src/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func TaskGetHandler(w http.ResponseWriter, r *http.Request) {
	var result *model.ApiResponse
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	vars := mux.Vars(r)
	stringId := vars["id"]
	id, err := strconv.Atoi(stringId)

	if stringId == "" {
		// TODO implement getting limit and offset from uri params
		result, err = taskModel.GetAll(500, 0)
		if err != nil {
			HandleApiError(w)
			return
		}
	} else if err == nil {
		result, err = taskModel.GetOne(id)
		if err != nil {
			HandleApiError(w)
			return
		}
	} else {
		result = &model.ApiResponse{JsonData: `{"error":"invalid id"}`, StatusCode: 400}
	}

	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}

func TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)
	var taskPostRequest model.TaskPostRequest
	err := decoder.Decode(&taskPostRequest)
	utils.LogIfError(err)

	result, err := taskModel.Post(taskPostRequest)
	if err != nil {
		HandleApiError(w)
		return
	}

	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}

func TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)
	var taskDeleteRequest model.TaskDeleteRequest
	err := decoder.Decode(&taskDeleteRequest)

	var result *model.ApiResponse
	if err == nil {
		result, err = taskModel.Delete(taskDeleteRequest)
		if err != nil {
			HandleApiError(w)
			return
		}
	} else {
		result = &model.ApiResponse{JsonData: `{"error":"invalid request body"}`, StatusCode: 422}
	}

	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}
