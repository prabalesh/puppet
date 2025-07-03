package model

import "time"

type InstallationJob struct {
	ID         int
	LanguageID int
	Install    bool
	Status     string    // "pending", "running", "done", "failed"
	Error      *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
