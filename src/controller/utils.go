package controller

import (
	"net/http"
)

func HandleApiError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
