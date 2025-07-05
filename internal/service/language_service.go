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
	logger *slog.Logger
}

func NewLanguageService(repo repository.LanguageRepository, logger *slog.Logger) *LanguageService {
	return &LanguageService{repo: repo, logger: logger}
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
	s.logger.Info("Updating installation", "id", id, "install", install)

	language, err := s.repo.GetLanguageById(id)
	if err != nil {
		s.logger.Error("Failed to get language for update", "error", err)
		return err
	}
	
	if language.Installed && install {
		return fmt.Errorf("Language with ID %d is already installed", id)
	}
	if !language.Installed && !install {
		return fmt.Errorf("Language with ID %d is not installed", id)
	}

	var cmd *exec.Cmd
	if install {
		cmd = exec.Command("docker", "pull", language.ImageName)
	} else {
		cmd = exec.Command("docker", "rmi", language.ImageName)
	}

	if err := cmd.Run(); err != nil {
		s.logger.Error("Docker operation failed", "error", err)
		return fmt.Errorf("docker error: %v", err)
	}

	return s.repo.UpdateInstallationStatus(id, install)
}
