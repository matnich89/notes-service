package handler

import (
	"encoding/json"
	"net/http"
	"notes-service/internal/helper"
	"notes-service/internal/logger"
	"notes-service/internal/note"
)

type Handler struct {
	service *note.Service
	logger  *logger.Logger
}

func NewHandler(service *note.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) PostNote(w http.ResponseWriter, r *http.Request) {

	var note note.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		h.logger.OutputError("Could not decode body into note")
		helper.SendBadRequest(w, "Could not decode body into notes", err)
		return
	}
	newNote, err := h.service.PostNote(note)
	if err != nil {
		h.logger.OutputError("Could not save new note")
		helper.SendInternalServerError(w, "Could not save note", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(newNote)

}
