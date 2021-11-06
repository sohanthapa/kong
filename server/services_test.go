package server

import (
	"testing"
)

//NOTE: The below tests are just the skeleton of how we are going to get the functions in
//       server/server.go. All the tests below are going to fail because we have not mocked
//       the database calls which are called in the handle function.

// TestHandleGETServices tests the functionality for handleGETServices function
func TestHandleGETServices(t *testing.T) {
	// t.Run("success", func(t *testing.T) {
	// 	request, _ := http.NewRequest("GET", "/services", nil)
	// 	response := httptest.NewRecorder()
	// 	server := &Server{
	// 		Router: chi.NewRouter(),
	// 	}

	// 	//call to initialize server
	// 	server.Initialize()
	// 	server.Router.ServeHTTP(response, request)
	// })
}

// TestHandleGETService tests the functionality for handleGETService function
func TestHandleGETService(t *testing.T) {
	// t.Run("success", func(t *testing.T) {
	// 	request, _ := http.NewRequest("GET", "/services/ServiceID", nil)
	// 	response := httptest.NewRecorder()
	// 	server := &Server{
	// 		Router: chi.NewRouter(),
	// 	}

	// 	//call to initialize server
	// 	server.Initialize()
	// 	server.Router.ServeHTTP(response, request)
	// })
}
