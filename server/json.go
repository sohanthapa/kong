package server

import (
	"encoding/json"
	"net/http"
)

// writeJSON encode data as JSON and sets the correct content type.
func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(contentTypeHTTPHeader, contentTypeApplicationJSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
