package model

import (
	"errors"
	"net/http"
	"repos/task_manager/src/db"
	"repos/task_manager/src/entity"
)

type TaskModel struct{}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t TaskModel) GetOne(id int) (ApiResponse, error) {
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	var task entity.Task

	db.Where("id = ?", id).Find(&task)

	if task.ID == 0 {
		return &TaskGetOneResponse{
			CreateDTOFromTaskModel(&task),
			http.StatusNotFound,
		}, nil
	}

	return &TaskGetOneResponse{
		CreateDTOFromTaskModel(&task),
		http.StatusOK,
	}, nil
}

func (t TaskModel) GetAll(limit int, offset int) (ApiResponse, error) {
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	db.Select("*").Limit(limit).Offset(offset).Find(&tasks)

	var result []*entity.TaskDTO

	for _, task := range tasks {
		result = append(result, CreateDTOFromTaskModel(&task))
	}

	return &TaskGetAllResponse{
		Tasks:      result,
		StatusCode: http.StatusOK,
	}, nil
}

func (t TaskModel) Post(req PostRequest) (ApiResponse, error) {
	postReq, ok := req.(TaskPostRequest)
	if !ok {
		return nil, errors.New("invalid post request")
	}

	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	var task entity.Task
	task.Title = postReq.Title
	task.Description = postReq.Description
	db.Create(&task)

	return &TaskPostResponse{TaskPost: TaskPost{task.ID, task.CreatedAt}, StatusCode: http.StatusCreated}, nil
}

func (t TaskModel) Put(req PutRequest) (ApiResponse, error) {
	putReq, ok := req.(TaskPutRequest)
	if !ok {
		return &TaskPutResponse{StatusCode: http.StatusBadRequest}, nil
	}
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	var task entity.Task
	db.Where("id = ?", putReq.Id).Find(&task)
	if task.ID == 0 {
		return &TaskPutResponse{StatusCode: http.StatusNotFound}, nil
	}

	task.Title = putReq.Title
	task.Description = putReq.Description
	db.Save(&task)

	return &TaskPutResponse{
		TaskPut: TaskPut{
			Id:        task.ID,
			UpdatedAt: task.UpdatedAt,
		},
		StatusCode: http.StatusOK,
	}, nil
}

func (t TaskModel) Delete(req DeleteRequest) (ApiResponse, error) {
	delReq, ok := req.(TaskDeleteRequest)
	if !ok {
		return &TaskDeleteResponse{StatusCode: http.StatusBadRequest}, nil
	}
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	db.Find(&tasks, "id in ?", delReq.Ids)
	db.Delete(&tasks)
	return &TaskDeleteResponse{StatusCode: http.StatusOK}, nil
}
