package database

import (
	"gorm.io/gorm"
	"notes-service/internal/note"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&note.Note{})
	if err != nil {
		return err
	}
	return nil
}
