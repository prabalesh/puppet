package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabalesh/puppet/internal/config"
	"github.com/prabalesh/puppet/internal/db"
	"github.com/prabalesh/puppet/internal/logging"
	"github.com/prabalesh/puppet/internal/repository/postgres"
	"github.com/prabalesh/puppet/internal/worker"
)

func main() {
	cfg := config.Load()
	logger := logging.NewLogger(cfg.Env)

	// Graceful shutdown setup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	database, err := db.InitDB(cfg.DBUrl)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()
	logger.Info("Database connection established")

	jobRepo := postgres.NewJobRepository(database)
	langRepo := postgres.NewLanguagePostgres(database)

	logger.Info("Installation worker started...")

	// Worker loop
	for {
		select {
		case <-ctx.Done():
			logger.Info("Shutdown signal received. Exiting...")
			return

		default:
			err := worker.ProcessNextJob(ctx, jobRepo, langRepo, logger)
			if err != nil {
				logger.Error("Job processing error", "error" , err)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
