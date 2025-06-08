package server

import (
	"context"
	"fmt"
	"log"
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

func Start(cfg *config.Config) {
	// Connect to the database
	database, err := db.InitDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	// Init services and handlers
	languageService := service.NewLanguageService(database)
	languageHandler := handler.NewLanguageHandler(languageService)

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
		fmt.Printf("Server started at PORT: %s\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v\n", err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	fmt.Println("Server gracefully stopped")
}
