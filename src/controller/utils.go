package controller

import (
	"fmt"
	"net/http"
)

func ApiReturnInternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func ApiReturnSuccess(w http.ResponseWriter, jsonString string, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, jsonString)
}

func ApiReturnBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, `{"error":"bad request"}`)
}
