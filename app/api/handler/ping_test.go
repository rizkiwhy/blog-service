package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler_Ping(t *testing.T) {
	// Initialize a new Gin engine in testing mode
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Create an instance of your PingHandler
	handler := NewPingHandler()

	// Register the Ping handler on the /ping route
	r.GET("/ping", handler.Ping)

	// Create a request to the /ping route
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body is correct
	assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())
}
