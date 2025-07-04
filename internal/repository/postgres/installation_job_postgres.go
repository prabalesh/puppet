package postgres

import (
	"database/sql"
	"time"

	"github.com/prabalesh/puppet/internal/model"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) CreateJob(job model.InstallationJob) (int, error) {
	query := `
		INSERT INTO language_installation_jobs (language_id, action, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(query, job.LanguageID, job.Action, job.Status).Scan(&id)
	return id, err
}

func (r *JobRepository) GetNextPendingJob() (*model.InstallationJob, error) {
	query := `
		SELECT id, language_id, action, status, error, created_at, updated_at
		FROM language_installation_jobs
		WHERE status = 'pending'
		ORDER BY created_at
		LIMIT 1
		FOR UPDATE SKIP LOCKED
	`
	row := r.db.QueryRow(query)

	var job model.InstallationJob
	err := row.Scan(
		&job.ID, &job.LanguageID, &job.Action, &job.Status,
		&job.Error, &job.CreatedAt, &job.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &job, err
}

func (r *JobRepository) UpdateJobStatus(id int, status string, errorMsg *string) error {
	query := `
		UPDATE language_installation_jobs
		SET status = $1, error = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := r.db.Exec(query, status, errorMsg, time.Now(), id)
	return err
}

func (r *JobRepository) GetJobByID(id int) (*model.InstallationJob, error) {
	query := `
		SELECT id, language_id, action, status, error, created_at, updated_at
		FROM language_installation_jobs
		WHERE id = $1
	`
	row := r.db.QueryRow(query, id)

	var job model.InstallationJob
	err := row.Scan(
		&job.ID, &job.LanguageID, &job.Action, &job.Status,
		&job.Error, &job.CreatedAt, &job.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &job, err
}
