package service

import (
	"fmt"
	"log/slog"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

type LanguageService struct {
	repo    repository.LanguageRepository
	jobRepo repository.JobInstallationRepository
	logger  *slog.Logger
}

func NewLanguageService(repo repository.LanguageRepository, jobRepo repository.JobInstallationRepository, logger *slog.Logger) *LanguageService {
	return &LanguageService{repo: repo, jobRepo: jobRepo, logger: logger}
}

func (s *LanguageService) ListLanguages() ([]model.Language, error) {
	s.logger.Info("Listing all languages")
	return s.repo.ListLanguages()
}

func (s *LanguageService) AddLanguage(lang model.Language) error {
	s.logger.Info("Adding new language", "name", lang.Name, "version", lang.Version)
	return s.repo.AddLanguage(lang)
}

func (s *LanguageService) DeleteLanguage(id int) error {
	s.logger.Info("Deleting language", "id", id)

	lang, err := s.repo.GetLanguageById(id)
	if err != nil {
		s.logger.Error("Failed to fetch language", "id", id, "error", err)
		return fmt.Errorf("language not found: %w", err)
	}

	if lang.Installed {
		s.logger.Info("Language is installed. Queuing uninstall job before deletion.", "id", id)

		job := model.InstallationJob{
			LanguageID: id,
			Action:     "delete",
			Status:     "pending",
		}

		if _, err := s.jobRepo.CreateJob(job); err != nil {
			s.logger.Error("Failed to queue uninstall job", "id", id, "error", err)
			return fmt.Errorf("failed to queue uninstall job: %w", err)
		}

		return nil
	}

	if _, _, err := s.repo.DeleteLanguage(id); err != nil {
		s.logger.Error("Failed to delete language", "id", id, "error", err)
		return fmt.Errorf("failed to delete language: %w", err)
	}

	s.logger.Info("Language deleted immediately", "id", id)
	return nil
}

func (s *LanguageService) UpdateInstallation(id int, install bool) (int, error) {
	s.logger.Info("Queuing install/uninstall job", "id", id, "install", install)
	var job model.InstallationJob

	_, err := s.repo.GetLanguageById(id)
	if err != nil {
		return -1, err
	}

	var action string = "uninstall"
	if install {
		action = "install"
	}
	job = model.InstallationJob{
		LanguageID: id,
		Action:     action,
		Status:     "pending",
	}
	jobId, err := s.jobRepo.CreateJob(job)
	if err != nil {
		s.logger.Error("Failed to create job", "error", err)
		return -1, err
	}

	return jobId, nil
}
