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

func (s *Service) GetNote(ID uint) (Note, error) {
	var note Note
	if result := s.DB.First(&note, ID); result.Error != nil {
		return Note{}, result.Error
	}
	return note, nil
}

func (s *Service) PostNote(note Note) (Note, error) {
	if result := s.DB.Save(&note); result.Error != nil {
		return Note{}, result.Error
	}
	return note, nil
}
