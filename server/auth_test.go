package server

import (
	"kong/data"
	"kong/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

//TestRequirePermission tests the functionality for RequirePermission func
func TestRequirePermission(t *testing.T) {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.With(RequirePermission(models.PermGetService, "NOT FOUND")).Get("/services/ServiceID1", func(w http.ResponseWriter, r *http.Request) {
		})
		r.With(RequirePermission(models.PermGetService, data.User2.ID)).Get("/services/ServiceID3", func(w http.ResponseWriter, r *http.Request) {
		})
		r.With(RequirePermission(models.PermGetServiceList, data.User2.ID)).Get("/services", func(w http.ResponseWriter, r *http.Request) {
		})
	})
	t.Run("http 403 if user not found by ID", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/services/ServiceID1", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusForbidden, response.Code, response.Body.String())
	})

	t.Run("http 403 if insufficient permissions", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/services/ServiceID3", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusForbidden, response.Code, response.Body.String())
	})

	//we cannot test the success case because we need to mock the database call
}

//TestAuthorizeUser tests the functionality for AuthorizeUser
func TestAuthorizeUser(t *testing.T) {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(AuthorizeUser(data.User1))
		r.With(RequirePermission(models.PermGetService, data.User1.ID)).Get("/services", func(w http.ResponseWriter, r *http.Request) {
		})
	})
	t.Run("success - user found", func(t *testing.T) {
		r := chi.NewRouter()
		r.Group(func(r chi.Router) {
			r.Use(AuthorizeUser(data.User1))
			r.With(RequirePermission(models.PermGetService, data.User1.ID)).Get("/services", func(w http.ResponseWriter, r *http.Request) {
			})
		})
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/services", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code, response.Body.String())

	})

	t.Run("error - user found", func(t *testing.T) {
		r := chi.NewRouter()
		r.Group(func(r chi.Router) {
			r.Use(AuthorizeUser(models.User{}))
			r.With(RequirePermission(models.PermGetService, data.User1.ID)).Get("/services", func(w http.ResponseWriter, r *http.Request) {
			})
		})
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/services", nil)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusForbidden, response.Code, response.Body.String())

	})
}
