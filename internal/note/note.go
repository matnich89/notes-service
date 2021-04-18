package note

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type Note struct {
	gorm.Model
	Name  string
	Theme string
	Text  string
	Owner uint
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

type NoteService interface {
	GetNote(ID uint) (Note, error)
	//GetNotesByTheme(theme string, owner uint) ([]Note, error)
	//PostNote(note Note)
}

func (s *Service) GetNote(ID uint) (Note, error) {
	var note Note
	if result := s.DB.First(&note, ID); result.Error != nil {
		return Note{}, result.Error
	}
	return note, nil
}
