package service

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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

func (s *ExecutorService) RunCode(runCodeDto dto.ExecuteCodeRequest) (string, string, string, error) {
	language, err := s.languageRepo.GetLanguageById(runCodeDto.LanguageID)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to get language: %w", err)
	}

	tmpDir, err := os.MkdirTemp("", "code-*")
	if err != nil {
		return "", "", "", fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	codePath := filepath.Join(tmpDir, language.FileName)
	if err := os.WriteFile(codePath, []byte(runCodeDto.Code), 0644); err != nil {
		return "", "", "", fmt.Errorf("failed to write code file: %w", err)
	}

	containerTmpPath := "/tmp"
	compile := language.CompileCommand
	run := language.RunCommand

	var dockerCmd string
	if compile != "" {
		dockerCmd = fmt.Sprintf("%s && %s", compile, run)
	} else {
		dockerCmd = run
	}

	cmd := exec.Command("docker", "run", "--rm", "-i",
		"-v", fmt.Sprintf("%s:%s", tmpDir, containerTmpPath),
		"-w", containerTmpPath,
		"--network", "none",
		language.ImageName,
		"sh", "-c", dockerCmd,
	)

	cmd.Stdin = bytes.NewReader([]byte(runCodeDto.Stdin))

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = tmpDir

	start := time.Now()
	err = cmd.Run()
	duration := time.Since(start).Seconds()
	durationStr := fmt.Sprintf("%.3f", duration)

	// if err != nil {
	// 	return out.String(), stderr.String(), durationStr, fmt.Errorf("execution error: %s - %w", stderr.String(), err)
	// }

	return out.String(), stderr.String(), durationStr, nil
}
