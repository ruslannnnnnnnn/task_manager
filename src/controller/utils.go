package controller

import (
	"fmt"
	"net/http"
	"repos/task_manager/src/model"
)

func HandleApiError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// HandleApiDefaultBehaviour returns to client such responses as 200, 201, 400, 404, 422
func HandleApiDefaultBehaviour(w http.ResponseWriter, result *model.ApiResponse) {
	w.WriteHeader(result.StatusCode)
	fmt.Fprint(w, result.JsonData)
}
