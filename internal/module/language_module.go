package module

import (
	"database/sql"
	"log/slog"

	"github.com/prabalesh/puppet/internal/handler"
	sqliterepo "github.com/prabalesh/puppet/internal/repository/sqlite"
	"github.com/prabalesh/puppet/internal/service"
)

func InitLanguageModule(db *sql.DB, logger *slog.Logger) *handler.LanguageHandler {
	repo := sqliterepo.NewLanguageSQLite(db)

	langService := service.NewLanguageService(repo, logger)
	langHandler := handler.NewLanguageHandler(langService, logger)
	return langHandler
}
