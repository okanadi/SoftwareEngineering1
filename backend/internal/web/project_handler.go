package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(svc *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: svc,
	}
}

func (h *ProjectHandler) HandleCreateProject(w http.ResponseWriter, r *http.Request) {
	// 1. JSON parsen
	var input domain.CreateProjectDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ungültiges JSON Format", http.StatusBadRequest)
		return
	}

	// 2. Service aufrufen
	newID, err := h.service.CreateProject(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Antwort senden
	response := map[string]string{
		"message": "Projekt erfolgreich angelegt",
		"id":      newID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *ProjectHandler) HandleGetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) HandleGetProjectByID(w http.ResponseWriter, r *http.Request) {
	projectID := path.Base(r.URL.Path)

	project, err := h.service.GetProjectByID(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) HandleGetProjectByCustomerLastname(w http.ResponseWriter, r *http.Request) {
	lastname := path.Base(r.URL.Path)

	projects, err := h.service.GetProjectByCustomerLastname(r.Context(), lastname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) HandleGetProjectByAddress(w http.ResponseWriter, r *http.Request) {
	address := path.Base(r.URL.Path)

	projects, err := h.service.GetProjectByAddress(r.Context(), address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) HandleGetAllCustomerLastnames(w http.ResponseWriter, r *http.Request) {
	lastnames, err := h.service.GetAllCustomerLastnames(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lastnames)
}

func (h *ProjectHandler) HandleGetAllAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.service.GetAllAddresses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

func (h *ProjectHandler) HandleGetByManagerID(w http.ResponseWriter, r *http.Request) {
	managerID := chi.URLParam(r, "managerID")

	projects, err := h.service.GetByManagerID(r.Context(), managerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
