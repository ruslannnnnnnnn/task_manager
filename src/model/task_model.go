package model

import (
	"encoding/json"
	"log"
	"repos/task_manager/src/db"
	"time"
)

type Task struct {
	Id          int       `json:id`
	Title       string    `json:title`
	Description string    `json:description`
	CreatedAt   time.Time `json:created_at`
}

func Get() string {

	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	result, err := database.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var tasks []Task

	for result.Next() {
		var task Task
		if err := result.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err := result.Err(); err != nil {
		log.Fatal(err)
	}

	tasksJson, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	return string(tasksJson)

}
