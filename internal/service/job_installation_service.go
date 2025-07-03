package service

import (
	"fmt"
	"log/slog"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

type JobInstallationService struct {
	Repo repository.JobInstallationRepository
	logger *slog.Logger
}

func NewJobInstallationService(repo repository.JobInstallationRepository, logger *slog.Logger) *JobInstallationService {
	return &JobInstallationService{Repo: repo, logger: logger}
}

func (s *JobInstallationService) GetJobStatus(id int) (*model.InstallationJob, error) {
	job, err := s.Repo.GetJobByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %w", err)
	}
	if job == nil {
		return nil, fmt.Errorf("job with ID %d not found", id)
	}
	return job, nil
}
