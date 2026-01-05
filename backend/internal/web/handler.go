package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

// POST /api/v1/customer/login
func (h *ProjectHandler) HandleCustomerLogin(w http.ResponseWriter, r *http.Request) {
	var req domain.CustomerLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ungültiges JSON", http.StatusBadRequest)
		return
	}

	project, err := h.service.LoginCustomer(r.Context(), req)
	if err != nil {
		http.Error(w, "Projekt nicht gefunden oder Daten falsch", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// GET /api/v1/projects/{id}/steps
func (h *ProjectHandler) HandleGetSteps(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	steps, err := h.service.GetProjectSteps(r.Context(), id)
	if err != nil {
		http.Error(w, "Fehler beim Laden der Schritte", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}

// POST /api/v1/steps/{id}/update (Multipart Form)
func (h *ProjectHandler) HandleUpdateStep(w http.ResponseWriter, r *http.Request) {
	stepID := chi.URLParam(r, "id")

	// Max 10MB
	r.ParseMultipartForm(10 << 20)

	input := domain.UpdateStepInput{
		StepID:    stepID,
		NewStatus: r.FormValue("status"), // MUSS übereinstimmen mit DB Enum: in_arbeit, problem, ...
		Note:      r.FormValue("note"),
		// TODO: Hier später UserID aus JWT Token holen
		UserID: "", // Bleibt leer -> wird NULL in DB
	}

	// Datei lesen
	file, header, err := r.FormFile("photo")
	if err == nil {
		defer file.Close()
		input.File = file
		input.FileName = header.Filename
		input.ContentType = header.Header.Get("Content-Type")
	}

	if err := h.service.UpdateStepProgress(r.Context(), input); err != nil {
		http.Error(w, "Update fehlgeschlagen: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}
