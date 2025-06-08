package sqlite

import (
	"database/sql"
	"errors"
	"time"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

type LanguageSQLite struct {
	db *sql.DB
}

func NewLanguageSQLite(db *sql.DB) repository.LanguageRepository {
	return &LanguageSQLite{db: db}
}

func (r *LanguageSQLite) ListLanguages() ([]model.Language, error) {
	rows, err := r.db.Query("SELECT * FROM languages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var langs []model.Language
	for rows.Next() {
		var lang model.Language
		err := rows.Scan(&lang.ID, &lang.Name, &lang.Version, &lang.ImageName, &lang.Installed, &lang.CreatedAt, &lang.UpdatedAt)
		if err != nil {
			return nil, err
		}
		langs = append(langs, lang)
	}
	return langs, nil
}

func (r *LanguageSQLite) AddLanguage(lang model.Language) error {
	now := time.Now()
	_, err := r.db.Exec(`INSERT INTO languages 
		(name, version, image_name, installed, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		lang.Name, lang.Version, lang.ImageName, false, now, now)
	return err
}

func (r *LanguageSQLite) DeleteLanguage(id int) (string, bool, error) {
	var imageName string
	var installed bool
	err := r.db.QueryRow("SELECT image_name, installed FROM languages WHERE id = ?", id).Scan(&imageName, &installed)
	if err == sql.ErrNoRows {
		return "", false, errors.New("language not found")
	} else if err != nil {
		return "", false, err
	}

	_, err = r.db.Exec("DELETE FROM languages WHERE id = ?", id)
	if err != nil {
		return "", false, err
	}
	return imageName, installed, nil
}

func (r *LanguageSQLite) UpdateInstallationStatus(id int, installed bool) error {
	_, err := r.db.Exec("UPDATE languages SET installed = ?, updated_at = ? WHERE id = ?", installed, time.Now(), id)
	return err
}

func (r *LanguageSQLite) GetLanguageById(id int) (model.Language, error) {
	var language model.Language
	err := r.db.QueryRow("SELECT * FROM languages WHERE id = ?", id).Scan(&language.ID, &language.Name, &language.Version, &language.ImageName, &language.Installed, &language.CreatedAt, &language.UpdatedAt)
	if err == sql.ErrNoRows {
		return model.Language{}, err
	} else if err != nil {
		return model.Language{}, err
	}

	return language, nil
}
