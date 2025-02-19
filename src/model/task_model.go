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

// GetOne returns a task by its ID
// @Summary Get a task by ID
// @Description Returns a task by its identifier
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} entity.TaskDTO
// @Failure 404 {object} controller.TaskNotFoundResponse
// @Router /tasks.json/{id} [get]
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

// GetAll returns a list of tasks with pagination
// @Summary Get a list of tasks
// @Description Returns a list of tasks with pagination
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {array} entity.TaskDTO
// @Router /tasks.json [get]
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

// Post creates a new task
// @Summary Create a task
// @Description Creates a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body TaskPostRequest true "Task data"
// @Success 201 {object} TaskPost
// @Failure 400 {object} controller.BadRequestResponse
// @Router /tasks.json [post]
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

// Put updates an existing task
// @Summary Update a task
// @Description Updates an existing task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body TaskPutRequest true "Task data"
// @Success 200 {object} TaskPut
// @Failure 400 {object} controller.BadRequestResponse
// @Failure 404 {object} controller.TaskNotFoundResponse
// @Router /tasks.json [put]
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

// Delete deletes tasks by their IDs
// @Summary Delete tasks
// @Description Deletes tasks by their IDs
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param ids body TaskDeleteRequest true "List of task IDs"
// @Success 200 {object} TaskDeleteResponse
// @Failure 400 {object} controller.BadRequestResponse
// @Router /tasks.json [delete]
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
	if len(tasks) > 0 {
		db.Delete(&tasks)
	}
	return &TaskDeleteResponse{StatusCode: http.StatusOK}, nil
}
