package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS languages (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			version TEXT NOT NULL,
			image_name TEXT UNIQUE NOT NULL,
			file_name TEXT NOT NULL,
			compile_command TEXT,
			run_command TEXT NOT NULL,
			installed BOOLEAN NOT NULL,
			created_at TIMESTAMPTZ NOT NULL,
			updated_at TIMESTAMPTZ NOT NULL
		);
		CREATE TABLE IF NOT EXISTS language_installation_jobs (
			id SERIAL PRIMARY KEY,
			language_id INTEGER NOT NULL,
			action TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending', -- pending | running | done | failed
			error TEXT,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
