package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func setupUserService() *user.UserService {
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

	user, err := service.Login("john_doe", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "john.doe@example.com", user.Email)
}

func TestLoginIncorrectPassword(t *testing.T) {
	service := setupUserService()

	_, err := service.SignUp("john_doe", "john.doe@example.com", "password123")
	assert.Nil(t, err)

	user, err := service.Login("john_doe", "wrongpassword")
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "incorrect password", err.Error())
}

func TestLoginUserNotFound(t *testing.T) {
	service := setupUserService()

	user, err := service.Login("unknown_user", "password123")
	assert.NotNil(t, err)
	assert.Nil(t, user)
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

	// verify password has changed
	loggedInUser, err := service.Login("john_doe", "newpassword123")
	assert.Nil(t, err)
	assert.Equal(t, loggedInUser.Email, loggedInUser.Email)

}

func TestResetPasswordInvalidToken(t *testing.T) {
	service := setupUserService()

	// Attempt to reset password with an invalid token
	err := service.ResetPassword("invalidtoken", "newpassword123")
	assert.NotNil(t, err)
	assert.Equal(t, "invalid or expired token", err.Error())
}
