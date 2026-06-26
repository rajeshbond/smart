package device

import (
	"encoding/json"
	"net/http"
	"strings"
)

type DeviceHandler struct {
	Service *DeviceService
}

func NewDeviceHandler(service *DeviceService) *DeviceHandler {
	return &DeviceHandler{Service: service}
}

func (h *DeviceHandler) RegisterDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Malformed JSON payload format"})
		return
	}

	// Execute down through our Service domain layer
	if err := h.Service.RegisterDeviceRequest(r.Context(), req); err != nil {
		if strings.Contains(err.Error(), "database transaction aborted") {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "Device name or MQTT Username already exists"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal hardware configuration failure"})
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Device successfully registered across database and broker engine metrics!",
	})
}
