package test

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func setupUserService() *user.UserService {

	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env.test file")
	}

	repo := &mocks.MockUserRepository{Users: make(map[string]*models.User)}
	return &user.UserService{Repo: repo}
}

func TestSignupSuccess(t *testing.T) {
	service := setupUserService()

	user, err := service.SignUp("john_doe", "john.doe@example.com", "password123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "john_doe", user.Username)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.NotEqual(t, "password123", user.PasswordHash)
}

func TestLoginSuccess(t *testing.T) {
	service := setupUserService()

	_, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)

	token, err := service.Login("john_doe", "password123")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// VERIFYING THE JWT TOKEN
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	assert.Nil(t, err)
	assert.Equal(t, "john_doe", claims.Subject)
}

func TestLoginIncorrectPassword(t *testing.T) {
	service := setupUserService()

	_, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)

	token, err := service.Login("john_doe", "wrongpassword")
	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "incorrect password", err.Error())
}

func TestLoginUserNotFound(t *testing.T) {
	service := setupUserService()

	token, err := service.Login("unknown_user", "password123")
	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())
}

func TestRequestPasswordReset(t *testing.T) {
	service := setupUserService()

	_, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)

	token, err := service.RequestPasswordReset("john.doe@example.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestResetPasswordSuccess(t *testing.T) {
	service := setupUserService()

	_, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)

	token, err := service.RequestPasswordReset("john.doe@example.com")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	err = service.ResetPassword(token, "newpassword123")
	assert.Nil(t, err)

	jwtToken, err := service.Login("john_doe", "newpassword123")
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtToken)

}

func TestResetPasswordInvalidToken(t *testing.T) {
	service := setupUserService()

	// Attempt to reset password with an invalid token
	err := service.ResetPassword("invalidtoken", "newpassword123")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid or expired token", err.Error())
}

func TestJWTTokenExpiration(t *testing.T) {
	jwtSecretKey := os.Getenv("JWT_SECRET")

	claims := &jwt.StandardClaims{
		Subject:   "john_doe",
		ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(), // Token expired 1 hour ago
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	assert.Nil(t, err)

	// Verify that the token is expired
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "token is expired"))
}

func TestInvalidJWTToken(t *testing.T) {
	invalidToken := "invalidTokenString"

	// Try to parse the invalid token
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(invalidToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	assert.NotNil(t, err)
}
