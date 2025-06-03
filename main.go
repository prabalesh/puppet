package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
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

	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("POST /api/languages", addLanguage)
	mux.HandleFunc("GET /api/languages", listLanguages)
	mux.HandleFunc("DELETE /api/languages/{id}", deleteLanguage)
	mux.HandleFunc("POST /api/languages/{id}/installations", installLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}/installations", uninstallLanguage)

	// server listenting
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

func deleteLanguage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var imageName string
	var installed bool
	err = db.QueryRow("SELECT image_name, installed FROM languages WHERE id = ?", id).Scan(&imageName, &installed)
	if err == sql.ErrNoRows {
		http.Error(w, "language not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if installed {
		cmd := exec.Command("docker", "rmi", imageName)
		err := cmd.Run()
		if err != nil {
			http.Error(w, fmt.Sprintf("docker uninstall failed: %v", err), http.StatusInternalServerError)
			return
		}
	}

	_, err = db.Exec("DELETE FROM languages WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func installLanguage(w http.ResponseWriter, r *http.Request) {
	doUpdateInstallationStatus(w, r, true)
}

func uninstallLanguage(w http.ResponseWriter, r *http.Request) {
	doUpdateInstallationStatus(w, r, false)
}

func doUpdateInstallationStatus(w http.ResponseWriter, r *http.Request, install bool) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var imageName string
	err = db.QueryRow("SELECT image_name FROM languages WHERE id = ?", id).Scan(&imageName)
	if err != nil {
		http.Error(w, "Language not found", http.StatusNotFound)
		return
	}

	var cmd *exec.Cmd
	if install {
		cmd = exec.Command("docker", "pull", imageName)
	} else {
		cmd = exec.Command("docker", "rmi", imageName)
	}

	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Docker error: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE languages SET installed = ?, updated_at = ? WHERE id = ?", install, time.Now(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
