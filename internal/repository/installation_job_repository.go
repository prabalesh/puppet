package repository

import "github.com/prabalesh/puppet/internal/model"

type JobInstallationRepository interface {
	CreateJob(job model.InstallationJob) (int, error)
	GetNextPendingJob() (*model.InstallationJob, error)
	UpdateJobStatus(id int, status string, errMsg *string) error
	GetJobByID(id int) (*model.InstallationJob, error)
}
