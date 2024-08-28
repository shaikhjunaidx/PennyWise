package user

import (
	"errors"

	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type UserService struct {
	Repo UserRepository
}

// In-memory map that stores active reset tokens
var passwordResetTokens = make(map[string]*PasswordResetToken)

// SignUp registers a new user with a hashed password
func (s *UserService) SignUp(username, email, password string) (*models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user based on username and password
func (s *UserService) Login(username, password string) (*models.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := ComparePasswords(user.PasswordHash, password); err != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

// RequestPasswordReset generates a password reset token for the user
func (s *UserService) RequestPasswordReset(email string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	token, err := GenerateResetToken()
	if err != nil {
		return "", err
	}

	StoreResetToken(token, user.Email)

	// Placeholder for sending the token via email in the future
	return token, nil
}

// ResetPassword allows the user to reset their password using a valid token
func (s *UserService) ResetPassword(token, newPassword string) error {
	resetToken, err := ValidateResetToken(token)
	if err != nil {
		return err
	}

	user, err := s.Repo.FindByEmail(resetToken.UserEmail)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	if err := s.Repo.Update(user); err != nil {
		return err
	}

	InvalidateResetToken(token)

	return nil
}
