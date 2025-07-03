package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/prabalesh/puppet/internal/service"
)

type JobInstallationHandler struct {
	Service *service.JobInstallationService
	Logger  *slog.Logger
}

func NewJobInstallationHandler(service *service.JobInstallationService, logger *slog.Logger) *JobInstallationHandler {
	return &JobInstallationHandler{Service: service, Logger: logger}
}

func (h *JobInstallationHandler) GetJobStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		RespondWithError(w, http.StatusBadRequest, "Job ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid job ID")
		return
	}

	job, err := h.Service.GetJobStatus(id)
	if err != nil {
		h.Logger.Error("Failed to get job status", "id", id, "error", err)
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, job)
}
