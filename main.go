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

	h := handler.New(database)

	// handlers
	mux.HandleFunc("POST /api/languages", h.AddLanguage)
	mux.HandleFunc("GET /api/languages", h.ListLanguages)
	mux.HandleFunc("DELETE /api/languages/{id}", h.DeleteLanguage)
	mux.HandleFunc("POST /api/languages/{id}/installations", h.InstallLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}/installations", h.UninstallLanguage)

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
