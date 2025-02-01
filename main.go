package main

import (
	"log"
	"net/http"
	"repos/task_manager/src/controller"
	"repos/task_manager/src/db"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	taskController := controller.NewTaskController()

	// Routes
	r.HandleFunc("/api/task.json/{id}", taskController.GetController).Methods("GET")
	r.HandleFunc("/api/task.json/", taskController.GetController).Methods("GET")
	r.HandleFunc("/api/task.json", taskController.GetController).Methods("GET")

	r.HandleFunc("/api/task.json", taskController.PostController).Methods("POST")

	r.HandleFunc("/api/task.json", taskController.DeleteController).Methods("DELETE")
	// Routes

	db.InitAutoMigrations()

	port := ":8080"
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
