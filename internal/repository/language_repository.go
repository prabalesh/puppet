package repository

import "github.com/prabalesh/puppet/internal/model"

type LanguageRepository interface {
	ListLanguages() ([]model.Language, error)
	AddLanguage(lang model.Language) error
	DeleteLanguage(id int) (imageName string, installed bool, err error)
	UpdateInstallationStatus(id int, installed bool) error
	GetLanguageById(id int) (model.Language, error)
}
