package module

import (
	"database/sql"
	"log/slog"

	"github.com/prabalesh/puppet/internal/handler"
	pgRepo "github.com/prabalesh/puppet/internal/repository/postgres"
	"github.com/prabalesh/puppet/internal/service"
)

func InitJobInstallationModule(db *sql.DB, logger *slog.Logger) *handler.JobInstallationHandler {
	repo := pgRepo.NewJobRepository(db)

	jobInstallationService := service.NewJobInstallationService(repo, logger)
	jobInstallationHandler := handler.NewJobInstallationHandler(jobInstallationService, logger)
	return jobInstallationHandler
}