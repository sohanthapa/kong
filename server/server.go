package server

import (
	"github.com/go-chi/chi"
)

// Server has database and router
type Server struct {
	Router *chi.Mux
}

// Initialize initalized the server based on the set configuration
func (s *Server) Initialize() {
	s.initRoutes()
}
