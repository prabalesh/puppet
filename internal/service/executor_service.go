package service

import (
	"log/slog"

	"github.com/prabalesh/puppet/internal/repository"
)

type ExecutorService struct {
	languageRepo repository.LanguageRepository
	logger       *slog.Logger
}

func NewExecutorService(langRepo repository.LanguageRepository, logger *slog.Logger) *ExecutorService {
	return &ExecutorService{languageRepo: langRepo, logger: logger}
}
