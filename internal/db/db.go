package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS languages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			version TEXT,
			image_name TEXT UNIQUE,
			installed BOOLEAN,
			created_at DATETIME,
			updated_at DATETIME
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
