package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "repos/task_manager/docs"
	"repos/task_manager/src/controller"
	"repos/task_manager/src/db"

	"github.com/gorilla/mux"
)

// @title task_manager API
// @description it's a simple rest api app that allows you to perform CRUD operations on tasks
// @host localhost:8080
// @BasePath /api
func main() {
	r := mux.NewRouter()

	taskController := controller.NewTaskController()

	// Routes
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	r.HandleFunc("/api/tasks.json/{id}", taskController.GetController).Methods("GET")
	r.HandleFunc("/api/tasks.json/", taskController.GetController).Methods("GET")
	r.HandleFunc("/api/tasks.json", taskController.GetController).Methods("GET")

	r.HandleFunc("/api/tasks.json", taskController.PostController).Methods("POST")

	r.HandleFunc("/api/tasks.json", taskController.PutController).Methods("PUT")

	r.HandleFunc("/api/tasks.json", taskController.DeleteController).Methods("DELETE")
	// Routes

	db.InitAutoMigrations()

	port := ":8080"
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
