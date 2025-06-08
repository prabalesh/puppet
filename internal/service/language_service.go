package service

import (
	"database/sql"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/prabalesh/puppet/internal/model"
)

type LanguageService struct {
	db *sql.DB
}

func NewLanguageService(db *sql.DB) *LanguageService {
	return &LanguageService{db: db}
}

func (s *LanguageService) ListLanguages() ([]model.Language, error) {
	rows, err := s.db.Query("SELECT * FROM languages")
	if err != nil {
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
			return nil, err
		}
		languages = append(languages, lang)
	}
	return languages, nil
}

func (s *LanguageService) AddLanguage(lang model.Language) error {
	currentTime := time.Now()
	_, err := s.db.Exec(
		"INSERT INTO languages (name, version, image_name, installed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		lang.Name,
		lang.Version,
		lang.ImageName,
		false,
		currentTime,
		currentTime)

	return err
}

func (s *LanguageService) DeleteLanguage(id int) error {
	var imageName string
	var installed bool
	err := s.db.QueryRow("SELECT image_name, installed FROM languages WHERE id = ?", id).Scan(&imageName, &installed)
	if err == sql.ErrNoRows {
		return errors.New("language not found")
	} else if err != nil {
		return err
	}

	if installed {
		cmd := exec.Command("docker", "rmi", imageName)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("docker uninstall failed: %v", err)
		}
	}

	_, err = s.db.Exec("DELETE FROM languages WHERE id = ?", id)
	return err
}

func (s *LanguageService) UpdateInstallation(id int, install bool) error {
	var imageName string
	err := s.db.QueryRow("SELECT image_name FROM languages WHERE id = ?", id).Scan(&imageName)
	if err == sql.ErrNoRows {
		return errors.New("language not found")
	} else if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if install {
		cmd = exec.Command("docker", "pull", imageName)
	} else {
		cmd = exec.Command("docker", "rmi", imageName)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker error: %v", err)
	}

	_, err = s.db.Exec("UPDATE languages SET installed = ?, updated_at = ? WHERE id = ?", install, time.Now(), id)
	return err
}
