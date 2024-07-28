package middleware

import (
	"app/constant"
	"app/pkg"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET_KEY", "NKy0VxAWfRW9ur0UE2DsSuc7eUJrU/Gdd7Qdk4YuImYEGCCvt2RkVAB8r4NWySuSjX2/ziPfE3A/Rwd5sJ0+uPMUsW/mlJA4Q8JTiUY783jxLdmZ5iG//qU/FQSuROqEbSjGVfPzfji+hc0A7S5Z+dE8hPo0DOW88VFdKwoUbzQ=")
	// Create a test router with the Authorization middleware
	router := gin.New()
	router.Use(Authorization())

	// Create a test endpoint that requires authorization
	router.GET("/protected", Authorization(), func(c *gin.Context) {
		_, oke := c.Request.Context().Value("user").(pkg.MetaToken)
		if !oke {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constant.ErrAuthorization})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": "valid_user"})
	})

	// Test case: Authorization header is missing
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test case: Invalid Authorization header format
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid_token")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test case: Invalid token type
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Basic abcdefg")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test case: Valid token
	validToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzNDk5MzIsImlkIjoiNjVhOTYxYTktZjY1Yy00NzBjLWE4YWUtZjYwZjRkOTgzZGFmIn0.IW4XYdtw7auE1KE84qW5VGCuXGT8IQ_PuXB3CNfPrTE"
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", validToken)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"user":"valid_user"`)
}
