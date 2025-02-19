package controller

import (
	"fmt"
	"net/http"
)

const BadRequestJsonString = `{"error":"bad request"}`

func ApiReturnInternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func ApiReturnResponse(w http.ResponseWriter, jsonString string, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, jsonString)
}

func ApiReturnBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, BadRequestJsonString)
}

// TODO start using not only for docs

type TaskNotFoundResponse struct {
	Error string `json:"error" example:"Task not found"`
}
type BadRequestResponse struct {
	Error string `json:"error" example:"bad request"`
}
