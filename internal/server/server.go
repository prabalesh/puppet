package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabalesh/puppet/internal/config"
	"github.com/prabalesh/puppet/internal/db"
	"github.com/prabalesh/puppet/internal/handler"
	"github.com/prabalesh/puppet/internal/middleware"
	"github.com/prabalesh/puppet/internal/service"
)

func Start(cfg *config.Config, logger *slog.Logger) {
	// Connect to the database
	logger.Info("Initializing database", "db", cfg.DBUrl)
	database, err := db.InitDB(cfg.DBUrl)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()
	logger.Info("Database connection established")

	// Init services and handlers
	languageService := service.NewLanguageService(database, logger)
	languageHandler := handler.NewLanguageHandler(languageService, logger)

	// HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/languages", languageHandler.AddLanguage)
	mux.HandleFunc("GET /api/languages", languageHandler.ListLanguages)
	mux.HandleFunc("DELETE /api/languages/{id}", languageHandler.DeleteLanguage)
	mux.HandleFunc("POST /api/languages/{id}/installations", languageHandler.InstallLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}/installations", languageHandler.UninstallLanguage)

	// HTTP server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: middleware.WithCORS(mux, cfg.AllowedOrigin),
	}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("Starting HTTP server", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server start failed", "error", err)
		}
	}()

	<-stop
	logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Forced shutdown failed", "error", err)
	}

	logger.Info("Server gracefully stopped")
}
