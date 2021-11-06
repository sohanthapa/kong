package server

import (
	"kong/data"
	"kong/models"
	"net/http"
)

// RequirePermission handles checking if user has proper permissions to access the handler
func RequirePermission(p string, userID string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			user := LoggedInUser(userID)

			if user == nil {
				http.Error(w, "user not found", http.StatusForbidden)
				return
			}
			if !user.HasPermission(p) {
				http.Error(w, "insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// LoggedInUser returns the info for the logged in user
// NOTE: for the simplicity of this exercise, I just passed in the userID to retrieve a value.
// This is just a small mock up for this exercise to demonstrate how we could do the authorization
// for the handler function
func LoggedInUser(userID string) *models.User {
	for _, user := range data.UserList {
		if user.ID == userID {
			return &user
		}
	}
	return nil

}
