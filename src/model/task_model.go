package model

import (
	"encoding/json"
	"gorm.io/gorm"
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

type TaskDeleteRequest struct {
	Ids []int `json:"ids"`
}

type TaskPostResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
type TaskPutResponse struct {
	Id        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskDeleteResponse struct {
	Id        uint           `json:"id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t TaskModel) GetOne(id int) *ApiResponse {
	db, err := db.InitDB()
	utils.LogIfError(err)
	var task entity.Task

	db.Where("id = ?", id).Find(&task)

	if task.ID == 0 {
		return &ApiResponse{
			`{"error":"task not found"}`,
			404,
		}
	}

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
	postReq, ok := req.(TaskPostRequest)
	if !ok {
		return &ApiResponse{`{"error": "bad request"}`, 400}
	}

	db, err := db.InitDB()
	utils.LogIfError(err)

	var task entity.Task
	task.Title = postReq.Title
	task.Description = postReq.Description
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

func (t TaskModel) Delete(req DeleteRequest) *ApiResponse {
	delReq, ok := req.(TaskDeleteRequest)
	if !ok {
		return &ApiResponse{`{"error": "bad request"}`, 400}
	}
	db, err := db.InitDB()
	utils.LogIfError(err)

	var tasks []entity.Task
	db.Find(&tasks, "id in ?", delReq.Ids)
	if len(tasks) == 0 {
		return &ApiResponse{
			`{"error":"tasks not found"}`, 404,
		}
	}

	db.Delete(&tasks)

	var result []TaskDeleteResponse
	for _, task := range tasks {
		result = append(result, TaskDeleteResponse{task.ID, task.DeletedAt})
	}

	resultJson, err := json.Marshal(result)
	utils.LogIfError(err)

	return &ApiResponse{
		string(resultJson),
		200,
	}
}
