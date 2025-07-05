package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/service"
)

type LanguageHandler struct {
	Service *service.LanguageService
	logger  *slog.Logger
}

func NewLanguageHandler(s *service.LanguageService, logger *slog.Logger) *LanguageHandler {
	return &LanguageHandler{Service: s, logger: logger}
}

func (h *LanguageHandler) ListLanguages(w http.ResponseWriter, r *http.Request) {
	languages, err := h.Service.ListLanguages()
	if err != nil {
		h.logger.Error("Failed to list languages", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	RespondWithJSON(w, http.StatusOK, languages)
}

func (h *LanguageHandler) AddLanguage(w http.ResponseWriter, r *http.Request) {
	var lang model.Language
	if err := JsonDecode(w, r, &lang); err != nil {
		h.logger.Warn("Invalid request payload", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.Service.AddLanguage(lang); err != nil {
		h.logger.Error("Failed to add language", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to add language")
		return
	}
	RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Language added successfully"})
}

func (h *LanguageHandler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Invalid ID format", "id", idStr)
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	if err := h.Service.DeleteLanguage(id); err != nil {
		h.logger.Error("Failed to delete language", "id", id, "error", err)
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Language deleted successfully"})
}

func (h *LanguageHandler) InstallLanguage(w http.ResponseWriter, r *http.Request) {
	h.doInstallation(w, r, true)
}

func (h *LanguageHandler) UninstallLanguage(w http.ResponseWriter, r *http.Request) {
	h.doInstallation(w, r, false)
}

func (h *LanguageHandler) doInstallation(w http.ResponseWriter, r *http.Request, install bool) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Invalid ID format", "id", idStr)
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	jobId, err := h.Service.UpdateInstallation(id, install)
	if err != nil {
		h.logger.Error("Failed to update installation state", "id", id, "install", install, "error", err)
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// msg := "Language installed successfully"
	// if !install {
	// 	msg = "Language uninstalled successfully"
	// }
	RespondWithJSON(w, http.StatusOK, map[string]int{"job_id": jobId})
}
