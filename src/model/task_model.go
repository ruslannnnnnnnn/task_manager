package model

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"repos/task_manager/src/db"
	"repos/task_manager/src/utils"
	"time"
)

// TODO configure the app to return status 500 if it fails to handle the request

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskInsertResult struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskInsertRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) GetAll(limit int, offset int) *ApiResponse {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stm, err := db.Prepare("SELECT * FROM tasks LIMIT $1 OFFSET $2")
	utils.LogIfError(err)
	result, err := stm.Query(limit, offset)
	utils.LogIfError(err)
	defer stm.Close()

	var tasks []Task

	for result.Next() {
		var task Task
		err := result.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt)
		utils.LogIfError(err)
		tasks = append(tasks, task)
	}

	err = result.Err()
	utils.LogIfError(err)

	tasksJson, err := json.Marshal(tasks)
	utils.LogIfError(err)

	return &ApiResponse{string(tasksJson), 200}

}

func (r *TaskRepository) GetOne(id int) *ApiResponse {
	db, err := db.InitDB()
	utils.LogIfError(err)
	defer db.Close()

	stm, err := db.Prepare("SELECT * FROM tasks WHERE id = $1 LIMIT 1")
	utils.LogIfError(err)
	defer stm.Close()

	result := stm.QueryRow(id)

	var task Task

	err = result.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt)
	if err == sql.ErrNoRows {
		return &ApiResponse{`{"error":"task not found"}`, 404}
	}
	utils.LogIfError(err)

	err = result.Err()
	utils.LogIfError(err)

	taskJson, err := json.Marshal(task)
	utils.LogIfError(err)

	return &ApiResponse{string(taskJson), 200}
}

func (r *TaskRepository) Post(requestBody io.ReadCloser) *ApiResponse {
	decoder := json.NewDecoder(requestBody)
	var task Task
	err := decoder.Decode(&task)
	utils.LogIfError(err)

	db, err := db.InitDB()
	utils.LogIfError(err)
	defer db.Close()

	stm, err := db.Prepare(
		`INSERT INTO tasks (title, description) 
		VALUES ($1, $2) 
		RETURNING id, created_at`,
	)
	utils.LogIfError(err)

	result := stm.QueryRow(task.Title, task.Description)
	var insertResult TaskInsertResult
	err = result.Scan(&insertResult.Id, &insertResult.CreatedAt)
	utils.LogIfError(err)

	var returnDataJson []byte

	returnDataJson, err = json.Marshal(insertResult)
	utils.LogIfError(err)

	return &ApiResponse{string(returnDataJson), 200}
}
