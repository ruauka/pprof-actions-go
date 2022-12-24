// Package handler -  Package handler.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Request - input message struct.
type Request struct {
	NumberOne int `json:"number_one"`
	NumberTwo int `json:"number_two"`
}

// HealthCheck - check health.
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok!!1"}) //nolint:errcheck,gosec
}

// Add - add two numbers.
func Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := map[string]int{
		"result": req.NumberOne + req.NumberTwo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
}
