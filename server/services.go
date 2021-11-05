package server

import (
	"encoding/json"
	"kong/models"
	"net/http"
)

func (s *Server) handleGETService(w http.ResponseWriter, r *http.Request) {
	var service models.Service

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&service)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, service, http.StatusOK)

}

func (s *Server) handleGETServices(w http.ResponseWriter, r *http.Request) {
	var service models.Service

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&service)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
