package storage

import (
	"database/sql"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

// Пример функции для добавления заметок
func (s *Storage) AddNote(userID int64, note string) error {
	_, err := s.DB.Exec("INSERT INTO notes (user_id, note) VALUES ($1, $2)", userID, note)
	return err
}
