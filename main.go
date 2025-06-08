package main

import (
	"github.com/prabalesh/puppet/internal/config"
	"github.com/prabalesh/puppet/internal/server"
)

func main() {
	cfg := config.Load()
	server.Start(cfg)
}
