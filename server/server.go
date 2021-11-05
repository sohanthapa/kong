package server

import (
	"github.com/go-chi/chi"
)

// Server has database and router
type Server struct {
	cfg          *Config
	Router       *chi.Mux
	maxBodyBytes int64 // default is zero; this must be set to value greather than 0 or the server will not accept any request bodies
}

// Initialize initalized the server based on the set configuration
func (s *Server) Initialize() {
	s.initRoutes()
}

//New initializes the Server struct with the database provided and a router
func New(cfg *Config) *Server {
	if cfg == nil {
		cfg = &Config{}
	}
	s := &Server{
		cfg: cfg,
	}

	return s
}
