package api

import (
	"encoding/json"
	"net/http"

	"harbory-backend/internal/domain"
)

type SystemHandler struct {
	service domain.SystemService
}

func NewSystemHandler(service domain.SystemService) *SystemHandler {
	return &SystemHandler{service: service}
}

func (h *SystemHandler) GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	info, err := h.service.GetInfo(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
