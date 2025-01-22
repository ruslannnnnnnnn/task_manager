package model

import (
	"database/sql"
	"encoding/json"
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

type TaskRepository struct{}

func NewTaskResitory() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) GetAll() string {

	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	result, err := database.Query("SELECT * FROM tasks LIMIT 500")
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

	return string(tasksJson)

}

func (r *TaskRepository) GetOne(id int) string {
	database, err := db.InitDB()
	utils.LogIfError(err)
	defer database.Close()

	stm, err := database.Prepare("SELECT * FROM tasks WHERE id = ? LIMIT 1")
	utils.LogIfError(err)
	defer stm.Close()

	result := stm.QueryRow(id)

	var task Task

	err = result.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt)
	if err == sql.ErrNoRows {
		return "{}"
	}
	utils.LogIfError(err)

	err = result.Err()
	utils.LogIfError(err)

	taskJson, err := json.Marshal(task)
	utils.LogIfError(err)

	return string(taskJson)
}
