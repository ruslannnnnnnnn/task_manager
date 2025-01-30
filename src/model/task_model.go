package model

import (
	"encoding/json"
	"repos/task_manager/src/db"
	"repos/task_manager/src/entity"
	"repos/task_manager/src/utils"
	"time"
)

// TODO configure the app to return status 500 if it fails to handle the request

type TaskModel struct{}

type TaskPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPutRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPostResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
type TaskPutResponse struct {
	Id        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t TaskModel) GetOne(id int) *ApiResponse {
	db, err := db.InitDB()
	utils.LogIfError(err)
	var task entity.Task

	db.Where("id = ?", id).Find(&task)

	jsonResult, err := json.Marshal(task)
	utils.LogIfError(err)

	return &ApiResponse{
		string(jsonResult),
		200,
	}
}

func (t TaskModel) GetAll(limit int, offset int) *ApiResponse {
	db, err := db.InitDB()
	utils.LogIfError(err)

	var tasks []entity.Task
	db.Select("*").Limit(limit).Offset(offset).Find(&tasks)

	result, err := json.Marshal(tasks)
	utils.LogIfError(err)
	return &ApiResponse{
		string(result),
		200,
	}
}

func (t TaskModel) Post(req PostRequest) *ApiResponse {
	db, err := db.InitDB()
	utils.LogIfError(err)

	taskReq, ok := req.(TaskPostRequest)
	if !ok {
		return &ApiResponse{`{"error": "invalid request"}`, 400}
	}

	var task entity.Task
	task.Title = taskReq.Title
	task.Description = taskReq.Description
	db.Create(&task)

	result := TaskPostResponse{
		task.ID,
		task.CreatedAt,
	}

	resultJson, err := json.Marshal(result)
	utils.LogIfError(err)

	return &ApiResponse{string(resultJson), 200}
}

func (t TaskModel) Put(req PutRequest) *ApiResponse {
	//TODO implement me
	panic("implement me")
	//db, err := db.InitDB()
	//utils.LogIfError(err)
}

func (t TaskModel) Delete(id int) *ApiResponse {
	//TODO implement me
	panic("implement me")
	//db, err := db.InitDB()
	//utils.LogIfError(err)
}
