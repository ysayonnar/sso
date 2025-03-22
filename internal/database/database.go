package database

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	Db *sql.DB
}

func Connect() (*Storage, error) {
	const op = "database.Connect()"

	pathToDatabase := filepath.Join("..", "..", "internal", "database", "sso.db")
	db, err := sql.Open("sqlite3", pathToDatabase)
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	var storage Storage
	storage.Db = db

	_, err = storage.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
    		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    		email TEXT UNIQUE NOT NULL,
    		password_hash TEXT NOT NULL,
    		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	return &storage, nil
}
