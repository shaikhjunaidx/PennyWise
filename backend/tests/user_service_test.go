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
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Fatalf("Error loading .env.test file")
	}

	repo := &mocks.MockUserRepository{Users: make(map[string]*models.User)}
	return &user.UserService{Repo: repo}
}

func createJWTToken(subject string, expiresAt time.Time) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: expiresAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func verifyJWTToken(t *testing.T, token string, expectedSubject string) {
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	assert.Nil(t, err)
	assert.Equal(t, expectedSubject, claims.Subject)
}

func TestSignupSuccess(t *testing.T) {
	service := setupUserService()

	user, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)
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
	verifyJWTToken(t, token, "john_doe")
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

	resetToken, err := service.RequestPasswordReset("john.doe@example.com")
	assert.Nil(t, err)

	err = service.ResetPassword(resetToken, "newpassword123")
	assert.Nil(t, err)

	token, err := service.Login("john_doe", "newpassword123")
	assert.Nil(t, err)
	verifyJWTToken(t, token, "john_doe")
}

func TestResetPasswordInvalidToken(t *testing.T) {
	service := setupUserService()

	err := service.ResetPassword("invalidtoken", "newpassword123")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid or expired token", err.Error())
}

func TestJWTTokenExpiration(t *testing.T) {
	expiredToken, err := createJWTToken("john_doe", time.Now().Add(-1*time.Hour))
	assert.Nil(t, err)

	_, err = jwt.ParseWithClaims(expiredToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "token is expired"))
}

func TestInvalidJWTToken(t *testing.T) {
	invalidToken := "invalidTokenString"

	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(invalidToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	assert.NotNil(t, err)
}
