package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Language struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	ImageName string    `json:"image_name"`
	Installed bool      `json:"installed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *sql.DB

func main() {
	// database connection
	var err error
	db, err = sql.Open("sqlite3", "file:storage/puppet.db?cache=shared&_foreign_keys=on")
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
	mux.HandleFunc("POST /api/languages", addLanguage)
	mux.HandleFunc("GET /api/languages", listLanguages)

	// handlers

	fmt.Println("Server started at http://localhost:8080/")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v\n", err.Error())
		return
	}
}

// handler functions
func addLanguage(w http.ResponseWriter, r *http.Request) {
	var lang Language
	if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	currentTime := time.Now()
	_, err := db.Exec(
		"INSERT INTO languages (name, version, image_name, installed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		lang.Name,
		lang.Version,
		lang.ImageName,
		false,
		currentTime,
		currentTime)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func listLanguages(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM languages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var languages []Language
	for rows.Next() {
		var lang Language

		err := rows.Scan(&lang.ID,
			&lang.Name,
			&lang.Version,
			&lang.ImageName,
			&lang.Installed,
			&lang.CreatedAt,
			&lang.UpdatedAt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		languages = append(languages, lang)
	}
	json.NewEncoder(w).Encode(languages)
}
