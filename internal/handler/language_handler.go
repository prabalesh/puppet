package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type Handler struct{ db *sql.DB }

func New(database *sql.DB) *Handler {
	return &Handler{db: database}
}

type Language struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	ImageName string    `json:"image_name"`
	Installed bool      `json:"installed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *Handler) AddLanguage(w http.ResponseWriter, r *http.Request) {
	var lang Language
	if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	currentTime := time.Now()
	_, err := h.db.Exec(
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

func (h *Handler) ListLanguages(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT * FROM languages")
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

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var imageName string
	var installed bool
	err = h.db.QueryRow("SELECT image_name, installed FROM languages WHERE id = ?", id).Scan(&imageName, &installed)
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

	_, err = h.db.Exec("DELETE FROM languages WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) InstallLanguage(w http.ResponseWriter, r *http.Request) {
	h.doUpdateInstallationStatus(w, r, true)
}

func (h *Handler) UninstallLanguage(w http.ResponseWriter, r *http.Request) {
	h.doUpdateInstallationStatus(w, r, false)
}

func (h *Handler) doUpdateInstallationStatus(w http.ResponseWriter, r *http.Request, install bool) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var imageName string
	err = h.db.QueryRow("SELECT image_name FROM languages WHERE id = ?", id).Scan(&imageName)
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

	_, err = h.db.Exec("UPDATE languages SET installed = ?, updated_at = ? WHERE id = ?", install, time.Now(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
