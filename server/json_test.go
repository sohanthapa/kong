package server

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestWriteJSON tests the functionality for WriteJSON function
func TestWriteJSON(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		response := httptest.NewRecorder()
		s := "Testing"
		WriteJSON(response, s, 200)
		responseBody := strings.TrimSuffix(response.Body.String(), "\n")
		responseBody = strings.Trim(responseBody, "\"")
		assert.Equal(t, s, responseBody)
		assert.Equal(t, 200, response.Code)

	})
}
