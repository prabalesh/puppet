package module

import (
	"database/sql"
	"log/slog"

	"github.com/prabalesh/puppet/internal/handler"
	"github.com/prabalesh/puppet/internal/service"
)

func InitLanguageModule(db *sql.DB, logger *slog.Logger) *handler.LanguageHandler {
	langService := service.NewLanguageService(db, logger)
	langHandler := handler.NewLanguageHandler(langService, logger)
	return langHandler
}
