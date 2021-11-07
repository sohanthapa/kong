package server

import (
	"kong/data"
	"kong/models"

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

		//RequirePermission is the middleware which checks the authorization for each
		// handler func
		r.Use(AuthorizeUser(data.User1)) //checks for user authentication
		r.With(RequirePermission(models.PermGetService, data.User1.ID)).Get(apiServiceID, s.handleGETService)
		r.With(RequirePermission(models.PermGetServiceList, data.User2.ID)).Get(apiServices, s.handleGETServices)
	})

}
