package handler

import (
	"encoding/json"
	"net/http"
	"notes-service/internal/logger"
	"notes-service/internal/note"
)

type Handler struct {
	service *note.Service
	logger  *logger.Logger
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service *note.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) PostNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var note note.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {

	}

}
