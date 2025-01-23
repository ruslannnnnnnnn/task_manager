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

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type InsertResult struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) GetAll() (resultJson string, statusCode int) {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Query("SELECT * FROM tasks LIMIT 500")
	utils.LogIfError(err)
	defer result.Close()

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

	return string(tasksJson), 200

}

func (r *TaskRepository) GetOne(id int) (resultJson string, statusCode int) {
	db, err := db.InitDB()
	utils.LogIfError(err)
	defer db.Close()

	stm, err := db.Prepare("SELECT * FROM tasks WHERE id = ? LIMIT 1")
	utils.LogIfError(err)
	defer stm.Close()

	result := stm.QueryRow(id)

	var task Task

	err = result.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt)
	if err == sql.ErrNoRows {
		return `{"error":"task not found"}`, 404
	}
	utils.LogIfError(err)

	err = result.Err()
	utils.LogIfError(err)

	taskJson, err := json.Marshal(task)
	utils.LogIfError(err)

	return string(taskJson), 200
}

func (r *TaskRepository) Post(requestBody io.ReadCloser) (resultJson string, statusCode int) {
	decoder := json.NewDecoder(requestBody)
	var task Task
	err := decoder.Decode(&task)
	utils.LogIfError(err)

	db, err := db.InitDB()
	utils.LogIfError(err)
	defer db.Close()

	stm, err := db.Prepare(
		`INSERT INTO tasks (title, description) 
		VALUES (?, ?) 
		RETURNING id, created_at`,
	)
	utils.LogIfError(err)

	result := stm.QueryRow(task.Title, task.Description)
	var insertResult InsertResult
	err = result.Scan(&insertResult.Id, &insertResult.CreatedAt)
	utils.LogIfError(err)

	var returnDataJson []byte

	returnDataJson, err = json.Marshal(insertResult)
	utils.LogIfError(err)

	return string(returnDataJson), 200
}
