package main

import (
	"log"
	"net/http"
	"repos/task_manager/src/controller"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/task.json/{id}", controller.TaskHandler).Methods("GET")
	r.HandleFunc("/api/task.json", controller.TaskHandler).Methods("GET")
	// Routes

	port := ":8080"
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
