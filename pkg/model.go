package pkg

import "notes-service/internal/note"

type NoteDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Theme string `json:"theme"`
	Text  string `json:"text"`
	Owner uint   `json:"owner"`
}

func NewNoteDto(note *note.Note) *NoteDto {
	return &NoteDto{
		ID:    note.Model.ID,
		Name:  note.Name,
		Theme: note.Theme,
		Text:  note.Text,
		Owner: 0, // Hardcoded until user service is up and running
	}
}
