package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"
	"path"
)

type ProjectStepHandler struct {
	service *service.ProjectStepService
}

func NewProjectStepHandler(srv *service.ProjectStepService) *ProjectStepHandler {
	return &ProjectStepHandler{
		service: srv,
	}
}

func (h *ProjectStepHandler) HandleCreateProjectStep(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateProjectStepDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ungültiges JSON Format", http.StatusBadRequest)
		return
	}

	newID, err := h.service.CreateProjectStep(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Projekt-Schritt erfolgreich angelegt",
		"id":      newID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *ProjectStepHandler) HandleGetProjectSteps(w http.ResponseWriter, r *http.Request) {
	projectID := path.Base(r.URL.Path)

	steps, err := h.service.GetProjectSteps(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}
