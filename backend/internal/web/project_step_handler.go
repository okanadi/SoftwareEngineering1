package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	projectID := chi.URLParam(r, "projectID")

	steps, err := h.service.GetProjectSteps(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}

func (h *ProjectStepHandler) HandleGetProjectStepByID(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	stepID := chi.URLParam(r, "stepID")

	step, err := h.service.GetProjectStepByID(r.Context(), projectID, stepID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(step)
}

func (h *ProjectStepHandler) HandleUpdateStepProgress(w http.ResponseWriter, r *http.Request) {
	// Multipart-Form parsen (z.B. max 10MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Daten aus dem Formular extrahieren
	input := domain.UpdateProjectStepDTO{
		StepId:    chi.URLParam(r, "stepID"),
		UserId:    r.FormValue("user_id"), // Später via JWT
		NewStatus: r.FormValue("new_status"),
		Note:      r.FormValue("note"),
	}

	// Datei extrahieren (Feldname im Frontend: "photo")
	file, header, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()
		input.File = file
		input.FileName = header.Filename
		input.FileContentType = header.Header.Get("Content-Type")
	}

	// Service aufrufen
	if err := h.service.UpdateStepProgress(r.Context(), input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Erfolgreich aktualisiert"})
}
