package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"repos/task_manager/src/db"
	"repos/task_manager/src/entity"
	"time"
)

type TaskModel struct{}

type TaskPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskPutRequest struct {
	Id          string `json:"id"`
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

func (t TaskModel) GetOne(id int) (*ApiResponse, error) {
	db, err := db.InitDB()
	if err != nil {
		return &ApiResponse{}, err
	}
	var task entity.Task

	db.Where("id = ?", id).Find(&task)

	if task.ID == 0 {
		return &ApiResponse{
			`{"error":"task not found"}`,
			http.StatusNotFound,
		}, nil
	}

	jsonResult, err := json.Marshal(task)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &ApiResponse{
		string(jsonResult),
		http.StatusOK,
	}, nil
}

func (t TaskModel) GetAll(limit int, offset int) (*ApiResponse, error) {
	db, err := db.InitDB()
	if err != nil {
		return &ApiResponse{}, err
	}

	var tasks []entity.Task
	db.Select("*").Limit(limit).Offset(offset).Find(&tasks)

	result, err := json.Marshal(tasks)
	if err != nil {
		return &ApiResponse{}, err
	}
	return &ApiResponse{
		string(result),
		http.StatusOK,
	}, nil
}

func (t TaskModel) Post(req PostRequest) (*ApiResponse, error) {
	postReq, ok := req.(TaskPostRequest)
	if !ok {
		return &ApiResponse{`{"error": "bad request"}`, http.StatusBadRequest}, nil
	}

	db, err := db.InitDB()
	if err != nil {
		return &ApiResponse{}, err
	}

	var task entity.Task
	task.Title = postReq.Title
	task.Description = postReq.Description
	db.Create(&task)

	result := TaskPostResponse{
		task.ID,
		task.CreatedAt,
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &ApiResponse{string(resultJson), http.StatusCreated}, nil
}

func (t TaskModel) Put(req PutRequest) (*ApiResponse, error) {
	putReq, ok := req.(TaskPutRequest)
	if !ok {
		return &ApiResponse{`{"error": "bad request"}`, http.StatusBadRequest}, nil
	}
	db, err := db.InitDB()
	if err != nil {
		return &ApiResponse{}, err
	}

	var task entity.Task
	db.Where("id = ?", putReq.Id).Find(&task)
	if task.ID == 0 {
		return &ApiResponse{
			`{"error":"tasks not found"}`, http.StatusNotFound,
		}, nil
	}

	task.Title = putReq.Title
	task.Description = putReq.Description
	db.Save(&task)

	jsonResult, err := json.Marshal(task)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &ApiResponse{JsonData: string(jsonResult), StatusCode: http.StatusOK}, nil

}

func (t TaskModel) Delete(req DeleteRequest) (*ApiResponse, error) {
	delReq, ok := req.(TaskDeleteRequest)
	if !ok {
		return &ApiResponse{`{"error": "bad request"}`, http.StatusBadRequest}, nil
	}
	db, err := db.InitDB()
	if err != nil {
		return &ApiResponse{}, err
	}

	var tasks []entity.Task
	db.Find(&tasks, "id in ?", delReq.Ids)
	if len(tasks) == 0 {
		return &ApiResponse{
			`{"error":"tasks not found"}`, http.StatusNotFound,
		}, nil
	}

	db.Delete(&tasks)

	var result []TaskDeleteResponse
	for _, task := range tasks {
		result = append(result, TaskDeleteResponse{task.ID, task.DeletedAt})
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return &ApiResponse{}, err
	}

	return &ApiResponse{
		string(resultJson),
		http.StatusOK,
	}, nil
}
