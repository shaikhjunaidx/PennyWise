package mocks

import (
	"errors"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
	Users  map[string]*models.User
	Emails map[string]*models.User
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	if m.Users == nil {
		m.Users = make(map[string]*models.User)
	}
	if m.Emails == nil {
		m.Emails = make(map[string]*models.User)
	}
	m.Users[user.Username] = user
	m.Emails[user.Email] = user
	return args.Error(0)
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	if user, exists := m.Users[username]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	if user, exists := m.Emails[email]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) Update(user *models.User) error {
	args := m.Called(user)
	if _, exists := m.Users[user.Username]; exists {
		m.Users[user.Username] = user
		m.Emails[user.Email] = user
		return args.Error(0)
	}
	return errors.New("user not found")
}

func (m *MockUserRepository) Delete(user *models.User) error {
	args := m.Called(user)
	if _, exists := m.Users[user.Username]; exists {
		delete(m.Users, user.Username)
		delete(m.Emails, user.Email)
		return args.Error(0)
	}
	return errors.New("user not found")
}
