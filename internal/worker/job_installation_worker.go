package worker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/repository"
)

func ProcessNextJob(
	ctx context.Context,
	jobRepo repository.JobInstallationRepository,
	langRepo repository.LanguageRepository,
	logger *slog.Logger,
) error {
	job, err := jobRepo.GetNextPendingJob()
	if err != nil {
		return fmt.Errorf("failed to fetch job: %w", err)
	}
	if job == nil {
		return nil
	}

	logger.Info("Picked job", "jobID", job.ID, "action", job.Action)

	if err := jobRepo.UpdateJobStatus(job.ID, "running", nil); err != nil {
		return fmt.Errorf("failed to mark job as running: %w", err)
	}

	return executeJob(job, jobRepo, langRepo, logger)
}

func executeJob(
	job *model.InstallationJob,
	jobRepo repository.JobInstallationRepository,
	langRepo repository.LanguageRepository,
	logger *slog.Logger,
) error {
	lang, err := langRepo.GetLanguageById(job.LanguageID)
	if err != nil {
		msg := fmt.Sprintf("language fetch error: %v", err)
		jobRepo.UpdateJobStatus(job.ID, "failed", &msg)
		return errors.New(msg)
	}

	cmd := prepareDockerCommand(job.Action, lang.ImageName)
	logger.Info("Executing Docker command", "jobID", job.ID, "command", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		msg := fmt.Sprintf("docker command failed: %v - %s", err, string(output))
		jobRepo.UpdateJobStatus(job.ID, "failed", &msg)
		return errors.New(msg)
	}

	// Update language install status in DB
	if job.Action == "delete" {
		langRepo.DeleteLanguage(job.LanguageID)
	} else {
		var install bool = false
		if job.Action == "install" {
			install = true
		}
		if err := langRepo.UpdateInstallationStatus(lang.ID, install); err != nil {
			msg := fmt.Sprintf("language status update failed: %v", err)
			jobRepo.UpdateJobStatus(job.ID, "failed", &msg)
			return errors.New(msg)
		}
	}
	if err := jobRepo.UpdateJobStatus(job.ID, "done", nil); err != nil {
		return fmt.Errorf("job status final update failed: %w", err)
	}

	logger.Info("Job completed successfully", "jobID", job.ID)
	return nil
}

func prepareDockerCommand(action string, image string) *exec.Cmd {
	if action == "install" {
		return exec.Command("docker", "pull", image)
	}
	return exec.Command("docker", "rmi", image)
}
