package web

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	// 1. JSON parsen
	var input domain.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Ungültiges JSON Format", http.StatusBadRequest)
		return
	}

	// 2. Service aufrufen
	newID, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Antwort senden
	response := map[string]string{
		"message": "User erfolgreich angelegt",
		"id":      newID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
