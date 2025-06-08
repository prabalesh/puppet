package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabalesh/puppet/internal/db"
	"github.com/prabalesh/puppet/internal/handler"
	"github.com/prabalesh/puppet/internal/service"
)

var database *sql.DB

func main() {
	// database connection
	var err error
	database, err = db.InitDB("file:storage/puppet.db?cache=shared&_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	defer database.Close()

	mux := http.NewServeMux()

	languageService := service.NewLanguageService(database)
	languageHandler := handler.NewLanguageHandler(languageService)

	// handlers
	mux.HandleFunc("POST /api/languages", languageHandler.AddLanguage)
	mux.HandleFunc("GET /api/languages", languageHandler.ListLanguages)
	mux.HandleFunc("DELETE /api/languages/{id}", languageHandler.DeleteLanguage)
	mux.HandleFunc("POST /api/languages/{id}/installations", languageHandler.InstallLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}/installations", languageHandler.UninstallLanguage)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to listen for interrupt or terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server started at http://localhost:8080/")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v\n", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	fmt.Println("\nShutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	fmt.Println("Server gracefully stopped")

}
