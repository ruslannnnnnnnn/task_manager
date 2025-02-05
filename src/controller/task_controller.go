package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "gorm.io/gorm"
	"net/http"
	"net/url"
	"repos/task_manager/src/model"
	"strconv"

	"github.com/gorilla/mux"
)

const QueryPageParamName = "page"
const QueryLimitParamName = "limit"

type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

// GetController Handles GET requests for tasks
// @Summary Get tasks
// @Description Get tasks or single task by ID, also has pagination
// @Tags tasks
// @Produce json
// @Param id path int false "task ID"
// @Param page query int false "number of the page" default(1)
// @Param limit query int false "max amount of tasks" default(1000)
// @Success 200 {object} []entity.TaskForDocs{}|entity.TaskForDocs{} "task or list of tasks"
// @Failure 404 {string} task not found
// @Router /api/task.json/{id} [get]
// @Router /api/task.json [get]
func (This *TaskController) GetController(w http.ResponseWriter, r *http.Request) {
	var result *model.ApiResponse
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	vars := mux.Vars(r)
	stringId := vars["id"]
	id, AtoiErr := strconv.Atoi(stringId)

	if stringId == "" {
		// TODO filters for updated_at and created_at columns
		limit, offset, err := This.ParseGetParams(r)
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
	if err != nil {
		HandleApiError(w)
		return
	}

	result, err := taskModel.Post(taskPostRequest)
	if err != nil {
		HandleApiError(w)
		return
	}

	HandleApiDefaultBehaviour(w, result)
}

func (*TaskController) PutController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskModel := model.NewTaskModel()

	decoder := json.NewDecoder(r.Body)

	var taskPutRequest model.TaskPutRequest
	err := decoder.Decode(&taskPutRequest)
	if err != nil {
		HandleApiError(w)
		return
	}

	result, err := taskModel.Put(taskPutRequest)
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

func (*TaskController) ParseGetParams(r *http.Request) (limit int, offset int, err error) {
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
