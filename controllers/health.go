package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rochdi/fizz-buzz/services/health"
)

// HealthController defines the health controller
type HealthController struct {
	service health.Service
}

// NewHealthController return a new instance of HealthController
func NewHealthController(svc health.Service) *HealthController {
	return &HealthController{svc}
}

// Health an sanity check endpoint, to be used for monitoring
func (h *HealthController) Health(rw http.ResponseWriter, r *http.Request) {
	if err := h.service.Health(); err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode("alive")
}

// Stats allows users to know what the most frequent request has been
func (h *HealthController) Stats(rw http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetStats()
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(stats)
}
