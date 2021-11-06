package server

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server has database and router
type Server struct {
	Router     *chi.Mux
	MongoStore *mongo.Database
}

// Initialize initalized the server based on the set configuration
func (s *Server) Initialize() {
	s.initRoutes()
}
