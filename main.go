package main

import (
	"log"
	"net/http"
	"repos/task_manager/src/controller"
)

func main() {

	// Routes
	http.HandleFunc("/api/task.json", controller.TaskHandler)
	// Routes

	port := ":8080"
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
