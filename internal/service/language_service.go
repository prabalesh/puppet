package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"time"

	"github.com/prabalesh/puppet/internal/model"
)

type LanguageService struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewLanguageService(db *sql.DB, logger *slog.Logger) *LanguageService {
	return &LanguageService{db: db, logger: logger}
}

func (s *LanguageService) ListLanguages() ([]model.Language, error) {
	s.logger.Info("Listing all languages")
	rows, err := s.db.Query("SELECT * FROM languages")
	if err != nil {
		s.logger.Error("Query failed", "error", err)
		return nil, err
	}
	defer rows.Close()

	var languages []model.Language
	for rows.Next() {
		var lang model.Language

		err := rows.Scan(&lang.ID,
			&lang.Name,
			&lang.Version,
			&lang.ImageName,
			&lang.Installed,
			&lang.CreatedAt,
			&lang.UpdatedAt)

		if err != nil {
			s.logger.Error("Failed to scan row", "error", err)
			return nil, err
		}
		languages = append(languages, lang)
	}
	return languages, nil
}

func (s *LanguageService) AddLanguage(lang model.Language) error {
	s.logger.Info("Adding new language", "name", lang.Name, "version", lang.Version)
	currentTime := time.Now()
	_, err := s.db.Exec(
		"INSERT INTO languages (name, version, image_name, installed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		lang.Name,
		lang.Version,
		lang.ImageName,
		false,
		currentTime,
		currentTime)

	if err != nil {
		s.logger.Error("Insert failed", "error", err)
	}

	return err
}

func (s *LanguageService) DeleteLanguage(id int) error {
	s.logger.Info("Deleting language", "id", id)

	var imageName string
	var installed bool
	err := s.db.QueryRow("SELECT image_name, installed FROM languages WHERE id = ?", id).Scan(&imageName, &installed)
	if err == sql.ErrNoRows {
		s.logger.Warn("Language not found", "id", id)
		return errors.New("language not found")
	} else if err != nil {
		s.logger.Error("Select failed", "error", err)
		return err
	}

	if installed {
		s.logger.Info("Uninstalling Docker image", "image", imageName)
		cmd := exec.Command("docker", "rmi", imageName)
		if err := cmd.Run(); err != nil {
			s.logger.Error("Docker uninstall failed", "error", err)
			return fmt.Errorf("docker uninstall failed: %v", err)
		}
	}

	_, err = s.db.Exec("DELETE FROM languages WHERE id = ?", id)
	if err != nil {
		s.logger.Error("Delete failed", "error", err)
	}
	return err
}

func (s *LanguageService) UpdateInstallation(id int, install bool) error {
	s.logger.Info("Updating installation state", "id", id, "install", install)
	var imageName string
	err := s.db.QueryRow("SELECT image_name FROM languages WHERE id = ?", id).Scan(&imageName)
	if err == sql.ErrNoRows {
		s.logger.Warn("Language not found", "id", id)
		return errors.New("language not found")
	} else if err != nil {
		s.logger.Error("Select failed", "error", err)
		return err
	}

	var cmd *exec.Cmd
	if install {
		s.logger.Info("Pulling Docker image", "image", imageName)
		cmd = exec.Command("docker", "pull", imageName)
	} else {
		s.logger.Info("Removing Docker image", "image", imageName)
		cmd = exec.Command("docker", "rmi", imageName)
	}

	if err := cmd.Run(); err != nil {
		s.logger.Error("Docker command failed", "error", err)

		return fmt.Errorf("docker error: %v", err)
	}

	_, err = s.db.Exec("UPDATE languages SET installed = ?, updated_at = ? WHERE id = ?", install, time.Now(), id)
	if err != nil {
		s.logger.Error("Update failed", "error", err)
	}
	return err
}
