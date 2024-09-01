package mocks

import (
	"errors"

	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type MockUserRepository struct {
	Users map[string]*models.User
}

func (m *MockUserRepository) Create(user *models.User) error {
	if _, exists := m.Users[user.Username]; exists {
		return errors.New("username already exists")
	}
	m.Users[user.Username] = user
	return nil
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	for _, user := range m.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	if user, exists := m.Users[username]; exists {
		return user, nil
	}

	return nil, errors.New("user not found")
}

func (m *MockUserRepository) Update(user *models.User) error {
	if _, exists := m.Users[user.Username]; !exists {
		return errors.New("user not found")
	}
	m.Users[user.Username] = user
	return nil
}

func (m *MockUserRepository) Delete(user *models.User) error {
	if _, exists := m.Users[user.Username]; !exists {
		return errors.New("user not found")
	}
	delete(m.Users, user.Username)
	return nil
}
