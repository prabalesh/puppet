package service

import (
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

type LanguageService struct {
	repo   repository.LanguageRepository
	jobRepo repository.JobInstallationRepository
	logger *slog.Logger
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
	imageName, installed, err := s.repo.DeleteLanguage(id)
	if err != nil {
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

	return nil
}

func (s *LanguageService) UpdateInstallation(id int, install bool) error {
	s.logger.Info("Queuing install/uninstall job", "id", id, "install", install)

	_, err := s.repo.GetLanguageById(id)
	if err != nil {
		return err
	}

	job := model.InstallationJob{
		LanguageID: id,
		Install:    install,
		Status:     "pending",
	}
	_, err = s.jobRepo.CreateJob(job)
	if err != nil {
		s.logger.Error("Failed to create job", "error", err)
		return err
	}

	return nil
}
