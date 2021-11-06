package server

import (
	"kong/models"
	mongostore "kong/stores/mongodb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) handleGETService(w http.ResponseWriter, r *http.Request) {
	serviceID := chi.URLParam(r, "serviceID")

	service, err := mongostore.GetService(s.MongoStore, serviceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, service, http.StatusOK)
}

// handleGETServices handles the GET request for endpoint /services and returns all
// the services based on the query parameter
func (s *Server) handleGETServices(w http.ResponseWriter, r *http.Request) {
	// used for filtering and searching by name
	nameValue := r.URL.Query().Get("name")
	// using offset and limit from the url query parameter for pagination
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	sort := r.URL.Query().Get("sort") //by default asc

	//by default find all the services
	filter := bson.M{}
	if len(nameValue) > 0 {
		filter = bson.M{"name": nameValue}
	}

	//setting the url query parameter
	slQuery := models.ServiceListQuery{
		Limit:  int64(limit),
		Offset: int64(offset),
		Sort:   sort,
		Filter: filter,
	}

	services, err := mongostore.ListServices(s.MongoStore, slQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, services, http.StatusOK)
}
