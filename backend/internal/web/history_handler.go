package web

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HistoryHandler struct {
	service *service.HistoryService
}

func NewHistoryHandler(svc *service.HistoryService) *HistoryHandler {
	return &HistoryHandler{
		service: svc,
	}
}

func (h *HistoryHandler) HandleGetHistory(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")

	history, err := h.service.GetHistory(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
