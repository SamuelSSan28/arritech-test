package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRequestLoggingMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a test logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Create a new router with the middleware
	router := gin.New()
	router.Use(RequestLoggingMiddleware(logger))

	// Add a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	// Create a test request
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("User-Agent", "test-agent")
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCORSMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router with the middleware
	router := gin.New()
	router.Use(CORSMiddleware())

	// Add a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	// Test GET request
	t.Run("GET request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check CORS headers
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
		assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test OPTIONS request (preflight)
	t.Run("OPTIONS request", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check CORS headers
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
		assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	// Test POST request
	t.Run("POST request", func(t *testing.T) {
		// Add POST route
		router.POST("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "test"})
		})

		req, _ := http.NewRequest("POST", "/test", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Check CORS headers
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
		assert.Equal(t, "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestCORSMiddleware_Headers(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router with the middleware
	router := gin.New()
	router.Use(CORSMiddleware())

	// Add a test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	// Test with specific headers
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("X-CSRF-Token", "test-token")
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Cache-Control", "no-cache")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check that all expected headers are allowed
	allowedHeaders := w.Header().Get("Access-Control-Allow-Headers")
	assert.Contains(t, allowedHeaders, "X-CSRF-Token")
	assert.Contains(t, allowedHeaders, "Authorization")
	assert.Contains(t, allowedHeaders, "Cache-Control")
	assert.Contains(t, allowedHeaders, "Content-Type")
	assert.Contains(t, allowedHeaders, "X-Requested-With")
}
