package helper

import (
	"encoding/json"
	"net/http"
	"notes-service/internal/response"
)

func SendBadRequest(w http.ResponseWriter, message string, err error) {
	sendErrorResponse(w, message, err, 400)
}

func SendInternalServerError(w http.ResponseWriter, message string, err error) {
	sendErrorResponse(w, message, err, 500)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error, statusCode int) {
	w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response.Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
