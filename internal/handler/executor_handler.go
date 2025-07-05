package handler

import (
	"log/slog"
	"net/http"

	"github.com/prabalesh/puppet/internal/dto"
	"github.com/prabalesh/puppet/internal/service"
)

type ExecutorHandler struct {
	Server *service.ExecutorService
	logger *slog.Logger
}

func NewExecutorHandler(s *service.ExecutorService, logger *slog.Logger) *ExecutorHandler {
	return &ExecutorHandler{s, logger}
}

func (h *ExecutorHandler) RunCode(w http.ResponseWriter, r *http.Request) {
	var runCodeReqBody dto.ExecuteCodeRequest
	if err := JsonDecode(w, r, &runCodeReqBody); err != nil {
		h.logger.Warn("Invalid request payload", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	codeOutput, err := h.Server.RunCode(runCodeReqBody)
	if err != nil {
		h.logger.Error("Failed to run code", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to add language")
		return
	}

	res := map[string]string{"stdout": codeOutput}

	RespondWithJSON(w, http.StatusOK, res)
}
