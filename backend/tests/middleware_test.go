package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/shaikhjunaidx/pennywise-backend/internal/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	secretKey  = "my-super-secret-key"
	validToken string
)

func setup() {
	os.Setenv("JWT_SECRET", secretKey)

	// Generate a valid token
	claims := &jwt.StandardClaims{
		Subject:   "user_id",
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, _ = token.SignedString([]byte(secretKey))
}

func createRequest(token string) *http.Request {
	req, _ := http.NewRequest("GET", "/api/transactions", nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

func createHandler() http.Handler {
	return middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}))
}

func TestJWTMiddleware_ValidToken(t *testing.T) {
	setup()

	req := createRequest(validToken)
	rr := httptest.NewRecorder()
	handler := createHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Success", rr.Body.String())
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {
	setup()

	req := createRequest("invalidtoken")
	rr := httptest.NewRecorder()
	handler := createHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJWTMiddleware_MissingToken(t *testing.T) {
	setup()

	req := createRequest("")
	rr := httptest.NewRecorder()
	handler := createHandler()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
