package main

import (
	"github.com/prabalesh/puppet/internal/config"
	"github.com/prabalesh/puppet/internal/logging"
	"github.com/prabalesh/puppet/internal/server"
)

func main() {
	cfg := config.Load()
	logger := logging.NewLogger(cfg.Env)
	server.Start(cfg, logger)
}
