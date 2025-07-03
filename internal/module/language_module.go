package module

import (
	"database/sql"
	"log/slog"

	"github.com/prabalesh/puppet/internal/handler"
	pgRepo "github.com/prabalesh/puppet/internal/repository/postgres"
	"github.com/prabalesh/puppet/internal/service"
)

func InitLanguageModule(db *sql.DB, logger *slog.Logger) *handler.LanguageHandler {
	repo := pgRepo.NewLanguagePostgres(db)
	jobRepo := pgRepo.NewJobRepository(db)

	langService := service.NewLanguageService(repo, jobRepo, logger)
	langHandler := handler.NewLanguageHandler(langService, logger)
	return langHandler
}
