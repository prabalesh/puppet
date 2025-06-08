package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/prabalesh/puppet/internal/model"
	"github.com/prabalesh/puppet/internal/service"
)

type LanguageHandler struct {
	Service *service.LanguageService
}

func NewLanguageHandler(s *service.LanguageService) *LanguageHandler {
	return &LanguageHandler{Service: s}
}

func (h *LanguageHandler) ListLanguages(w http.ResponseWriter, r *http.Request) {
	langs, err := h.Service.ListLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(langs)
}

func (h *LanguageHandler) AddLanguage(w http.ResponseWriter, r *http.Request) {
	var lang model.Language
	if err := json.NewDecoder(r.Body).Decode(&lang); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.Service.AddLanguage(lang); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *LanguageHandler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteLanguage(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateInstallation(id, install); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
