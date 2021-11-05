package server

import (
	"github.com/go-chi/chi"
)

const (
	apiServiceID = "/services/{serviceID}"
	apiServices  = "/services"
)

// initRoutes initializes the routes for the router
func (s *Server) initRoutes() {

	// private routes
	s.Router.Group(func(r chi.Router) {
		r.With().Get(apiServiceID, s.handleGETService)
	})

}
