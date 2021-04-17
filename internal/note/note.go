package note

import "gorm.io/gorm"

type Service struct {
    DB *gorm.DB
}

type Note struct {
    gorm.Model
    name string
    theme string
    text string
    owner uint
}

type NoteService interface {
    GetNote(ID uint) (Note, error)
    GetNotesByTheme(theme string, owner uint)  ([]Note, error)
    PostNote(note Note)
}

func (s *Service) GetNote(ID uint) (Note, error) {
    var note Note
    if result := s.DB.First(&note, ID); result.Error  != nil {
        return Note{}, result.Error
    }
    return note, nil
}



