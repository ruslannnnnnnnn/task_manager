package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"repos/task_manager/src/model"
	"repos/task_manager/src/utils"
	"strconv"

	"github.com/gorilla/mux"
)

const QueryPageParamName = "page"
const QueryLimitParamName = "limit"

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (This *TaskController) GetController(w http.ResponseWriter, r *http.Request) {
	var result *model.ApiResponse
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	vars := mux.Vars(r)
	stringId := vars["id"]
	id, AtoiErr := strconv.Atoi(stringId)

	if stringId == "" {
		// TODO filters for updated_at and created_at columns
		limit, offset, err := This.ParseGetParams(w, r)
		if err != nil {
			switch err.Error() {
			case "500":
				HandleApiError(w)
				return
			case "400":
				HandleApiDefaultBehaviour(w, &model.ApiResponse{
					JsonData: fmt.Sprintf(
						`{"error":"invalid %s/%s param"}`,
						QueryPageParamName,
						QueryLimitParamName,
					),
					StatusCode: http.StatusBadRequest,
				})
				return
			case "max limit=1000":
				HandleApiDefaultBehaviour(w, &model.ApiResponse{
					JsonData:   `{"error":"limit can't be more than 1000"}`,
					StatusCode: http.StatusUnprocessableEntity,
				})
				return
			}
		}

		result, err = taskModel.GetAll(limit, (offset-1)*limit)
		if err != nil {
			HandleApiError(w)
			return
		}
	} else if AtoiErr == nil {
		result, AtoiErr = taskModel.GetOne(id)
		if AtoiErr != nil {
			HandleApiError(w)
			return
		}
	} else {
		result = &model.ApiResponse{JsonData: `{"error":"invalid id"}`, StatusCode: http.StatusBadRequest}
	}

	HandleApiDefaultBehaviour(w, result)
}

func (*TaskController) PostController(w http.ResponseWriter, r *http.Request) {
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

	HandleApiDefaultBehaviour(w, result)
}

func (*TaskController) DeleteController(w http.ResponseWriter, r *http.Request) {
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
		result = &model.ApiResponse{JsonData: `{"error":"bad request"}`, StatusCode: http.StatusBadRequest}
	}

	HandleApiDefaultBehaviour(w, result)
}

func (*TaskController) ParseGetParams(w http.ResponseWriter, r *http.Request) (limit int, offset int, err error) {
	limit = 1000

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return limit, offset, errors.New("500")
	}

	if query.Has(QueryPageParamName) {
		offset, err = strconv.Atoi(query.Get(QueryPageParamName))
		if err != nil {
			return limit, offset, errors.New("400")
		}
	}

	if query.Has(QueryLimitParamName) {
		limit, err = strconv.Atoi(query.Get(QueryLimitParamName))
		if err != nil {
			return limit, offset, errors.New("400")
		}
		if limit > 1000 {
			return limit, offset, errors.New("max limit=1000")
		}
	}

	return limit, offset, nil
}
