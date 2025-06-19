package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

type LanguagePostgres struct {
	db *sql.DB
}

func NewLanguagePostgres(db *sql.DB) repository.LanguageRepository {
	return &LanguagePostgres{db: db}
}

func (r *LanguagePostgres) ListLanguages() ([]model.Language, error) {
	rows, err := r.db.Query("SELECT id, name, version, image_name, file_name, compile_command, run_command, installed, created_at, updated_at FROM languages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var langs []model.Language
	for rows.Next() {
		var lang model.Language
		err := rows.Scan(&lang.ID, &lang.Name, &lang.Version, &lang.ImageName, &lang.FileName, &lang.CompileCommand, &lang.RunCommand, &lang.Installed, &lang.CreatedAt, &lang.UpdatedAt)
		if err != nil {
			return nil, err
		}
		langs = append(langs, lang)
	}
	return langs, nil
}

func (r *LanguagePostgres) AddLanguage(lang model.Language) error {
	now := time.Now()
	_, err := r.db.Exec(`
		INSERT INTO languages (name, version, image_name, file_name, compile_command, run_command, installed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, lang.Name, lang.Version, lang.ImageName, lang.FileName, lang.CompileCommand, lang.RunCommand, false, now, now)
	return err
}

func (r *LanguagePostgres) DeleteLanguage(id int) (string, bool, error) {
	var imageName string
	var installed bool
	err := r.db.QueryRow("SELECT image_name, installed FROM languages WHERE id = $1", id).Scan(&imageName, &installed)
	if err == sql.ErrNoRows {
		return "", false, errors.New("language not found")
	} else if err != nil {
		return "", false, err
	}

	_, err = r.db.Exec("DELETE FROM languages WHERE id = $1", id)
	if err != nil {
		return "", false, err
	}
	return imageName, installed, nil
}

func (r *LanguagePostgres) UpdateInstallationStatus(id int, installed bool) error {
	_, err := r.db.Exec("UPDATE languages SET installed = $1, updated_at = $2 WHERE id = $3", installed, time.Now(), id)
	return err
}

func (r *LanguagePostgres) GetLanguageById(id int) (model.Language, error) {
	var lang model.Language
	err := r.db.QueryRow("SELECT id, name, version, image_name, file_name, compile_command, run_command, installed, created_at, updated_at FROM languages WHERE id = $1", id).
		Scan(&lang.ID, &lang.Name, &lang.Version, &lang.ImageName, &lang.FileName, &lang.CompileCommand, &lang.RunCommand, &lang.Installed, &lang.CreatedAt, &lang.UpdatedAt)
	if err != nil {
		return model.Language{}, err
	}
	return lang, nil
}
