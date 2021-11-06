package server

import (
	"encoding/json"
	"net/http"
)

const (
	// HTTP header values
	contentTypeHTTPHeader      string = "Content-Type"
	contentTypeApplicationJSON string = "application/json"
)

// WriteJSON encode data as JSON and sets the correct content type.
func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(contentTypeHTTPHeader, contentTypeApplicationJSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
