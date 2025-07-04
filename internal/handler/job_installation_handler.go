package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/prabalesh/puppet/internal/service"
)

type JobInstallationHandler struct {
	Service *service.JobInstallationService
	logger  *slog.Logger
}

func NewJobInstallationHandler(service *service.JobInstallationService, logger *slog.Logger) *JobInstallationHandler {
	return &JobInstallationHandler{Service: service, logger: logger}
}

func (h *JobInstallationHandler) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Invalid ID format", "id", idStr)
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	job, err := h.Service.GetJobStatus(id)
	if err != nil {
		h.logger.Error("Failed to get job status", "id", id, "error", err)
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, job)
}
