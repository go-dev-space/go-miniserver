package handlers

import (
	"encoding/json"
	"net/http"
)

func (rh *RouteHandlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "wrong method type", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewEncoder(w).Encode(struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}{
		StatusCode: http.StatusOK,
		Message:    "it work's!",
	})

	if err != nil {
		http.Error(w, "uups, something went wrong!", http.StatusInternalServerError)
		return
	}
}
