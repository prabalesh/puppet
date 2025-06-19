package model

import "time"

type Language struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Version        string    `json:"version"`
	ImageName      string    `json:"image_name"`
	FileName       string    `json:"file_name"`
	CompileCommand string    `json:"compile_command"`
	RunCommand     string    `json:"run_command"`
	Installed      bool      `json:"installed"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
