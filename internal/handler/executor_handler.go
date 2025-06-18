package handler

import (
	"log/slog"

	"github.com/prabalesh/puppet/internal/service"
)

type ExecutorHandler struct {
	Server *service.ExecutorService
	logger *slog.Logger
}

func NewExecutorHandler(s *service.ExecutorService, logger *slog.Logger) *ExecutorHandler {
	return &ExecutorHandler{s, logger}
}
