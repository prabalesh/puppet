package module

import (
	"database/sql"
	"log/slog"

	"github.com/prabalesh/puppet/internal/handler"
	pgRepo "github.com/prabalesh/puppet/internal/repository/postgres"
	"github.com/prabalesh/puppet/internal/service"
)

func InitExecutorModule(db *sql.DB, logger *slog.Logger) *handler.ExecutorHandler {
	repo := pgRepo.NewLanguagePostgres(db)
	executorService := service.NewExecutorService(repo, logger)
	executorHandler := handler.NewExecutorHandler(executorService, logger)

	return executorHandler
}
