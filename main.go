package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// database connection
	var err error
	db, err = sql.Open("sqlite3", "file:storage/languages.db?cache=shared&_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS languages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			version TEXT,
			image_name TEXT UNIQUE,
			installed BOOLEAN,
			created_at DATETIME,
			updated_at DATETIME
		)`)
	if err != nil {
		log.Fatalf("failed to create the languages table%v", err.Error())
	}

	// server listenting
	mux := http.NewServeMux()

	// handlers

	fmt.Println("Server started at http://localhost:8080/")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v\n", err.Error())
		return
	}
}
