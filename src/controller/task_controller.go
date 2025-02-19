package controller

import (
	"encoding/json"
	_ "gorm.io/gorm"
	"net/http"
	"net/url"
	"repos/task_manager/src/model"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	QueryPageParamName       = "page"
	QueryLimitParamName      = "limit"
	TaskNotFoundJsonString   = `{"error": "Task not found"}`
	SuccessMessageJsonString = `{"message":"success"}`
)

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (This *TaskController) GetController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	vars := mux.Vars(r)
	stringId := vars["id"]
	id, AtoiErr := strconv.Atoi(stringId)

	if stringId == "" {
		limit, offset := This.ParseGetParams(r)
		getRes, err := taskModel.GetAll(limit, offset)
		if err != nil {
			ApiReturnInternalServerError(w)
			return
		}
		jsonString, err := json.Marshal(getRes.GetBody())
		if err != nil {
			ApiReturnInternalServerError(w)
			return
		}

		if getRes.GetStatus() >= 200 && getRes.GetStatus() <= 299 {
			ApiReturnResponse(w, string(jsonString), getRes.GetStatus())
		}
	} else if AtoiErr == nil {
		getRes, err := taskModel.GetOne(id)
		if err != nil {
			ApiReturnInternalServerError(w)
			return
		}
		jsonString, err := json.Marshal(getRes.GetBody())
		if getRes.GetStatus() >= 400 && getRes.GetStatus() <= 499 {
			ApiReturnResponse(w, TaskNotFoundJsonString, getRes.GetStatus())
			return
		}
		ApiReturnResponse(w, string(jsonString), getRes.GetStatus())
	} else {
		ApiReturnBadRequest(w)
	}
}

func (*TaskController) PostController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)
	var taskPostRequest model.TaskPostRequest
	err := decoder.Decode(&taskPostRequest)
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	result, err := taskModel.Post(taskPostRequest)
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	resultJson, err := json.Marshal(result.GetBody())
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	ApiReturnResponse(w, string(resultJson), result.GetStatus())
}

func (*TaskController) PutController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)

	var taskPutRequest model.TaskPutRequest
	err := decoder.Decode(&taskPutRequest)
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	result, err := taskModel.Put(taskPutRequest)
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	resultJson, err := json.Marshal(result.GetBody())
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	if result.GetStatus() >= 400 && result.GetStatus() <= 499 {
		ApiReturnResponse(w, TaskNotFoundJsonString, result.GetStatus())
	}

	ApiReturnResponse(w, string(resultJson), result.GetStatus())
}

func (*TaskController) DeleteController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)
	var taskDeleteRequest model.TaskDeleteRequest
	err := decoder.Decode(&taskDeleteRequest)
	if err != nil {
		ApiReturnBadRequest(w)
		return
	}
	result, err := taskModel.Delete(taskDeleteRequest)
	if err != nil {
		ApiReturnInternalServerError(w)
		return
	}

	ApiReturnResponse(w, SuccessMessageJsonString, result.GetStatus())
}

func (*TaskController) ParseGetParams(r *http.Request) (limit int, offset int) {
	limit = 1000
	offset = 1

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return limit, offset
	}

	if query.Has(QueryPageParamName) {
		offset, err = strconv.Atoi(query.Get(QueryPageParamName))
		if err != nil {
			return limit, offset
		}
	}

	if query.Has(QueryLimitParamName) {
		limit, err = strconv.Atoi(query.Get(QueryLimitParamName))
		if err != nil {
			return limit, offset
		}
		if limit > 1000 {
			return 1000, offset
		}
	}

	return limit, offset
}
