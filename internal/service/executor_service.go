package service

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/prabalesh/puppet/internal/dto"
	"github.com/prabalesh/puppet/internal/repository"
)

type ExecutorService struct {
	languageRepo repository.LanguageRepository
	logger       *slog.Logger
}

func NewExecutorService(langRepo repository.LanguageRepository, logger *slog.Logger) *ExecutorService {
	return &ExecutorService{languageRepo: langRepo, logger: logger}
}

func (s *ExecutorService) RunCode(runCodeDto dto.ExecuteCodeRequest) (string, error) {
	language, err := s.languageRepo.GetLanguageById(runCodeDto.LanguageID)
	if err != nil {
		return "", fmt.Errorf("failed to get language: %w", err)
	}

	dockerArgs := []string{
		"run", "--rm",
		language.ImageName,
	}

	var codeCmd []string
	switch language.Name {
	case "python":
		codeCmd = []string{"python3", "-c", runCodeDto.Code}
	case "javascript":
		codeCmd = []string{"node", "-e", runCodeDto.Code}
	default:
		return "", fmt.Errorf("unsupported language: %s", language.Name)
	}

	dockerArgs = append(dockerArgs, codeCmd...)

	cmd := exec.Command("docker", dockerArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("docker run failed: %s - %w", stderr.String(), err)
	}

	return out.String(), nil
}
