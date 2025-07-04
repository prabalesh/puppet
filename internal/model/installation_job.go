package model

import "time"

type InstallationJob struct {
	ID         int       `json:"id"`
	LanguageID int       `json:"language_id"`
	Action     string    `json:"action"`
	Status     string    `json:"status"`
	Error      *string   `json:"error"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
