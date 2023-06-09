package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"status": "up"}
	if err := enc.Encode(resp); err != nil {
		panic("failed to encode response")
	}
}
