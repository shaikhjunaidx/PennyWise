package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/testutils"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CRUDOperations(t *testing.T) {
	_, tx := testutils.SetupTestDB()
	defer tx.Rollback() // Rollback the transaction after the test

	repo := user.NewUserRepository(tx)

	// Test Create
	user := &models.User{
		Username:     "john_doe",
		Email:        "john.doe@example.com",
		PasswordHash: "hashed_password",
	}

	err := repo.Create(user)
	assert.Nil(t, err, "Failed to create user")

	// Test FindByUsername
	foundUser, err := repo.FindByUsername("john_doe")
	assert.Nil(t, err, "Failed to find user by username")
	assert.Equal(t, user.Email, foundUser.Email)

	// Test FindByEmail
	foundUser, err = repo.FindByEmail("john.doe@example.com")
	assert.Nil(t, err, "Failed to find user by email")
	assert.Equal(t, user.Username, foundUser.Username)

	// Test Update
	user.Username = "john_updated"
	err = repo.Update(user)
	assert.Nil(t, err, "Failed to update user")

	// Verify the update
	updatedUser, err := repo.FindByUsername("john_updated")
	assert.Nil(t, err, "Failed to find updated user by username")
	assert.Equal(t, "john_updated", updatedUser.Username)

	// Test Delete
	err = repo.Delete(user)
	assert.Nil(t, err, "Failed to delete user")

	// Verify the user is deleted
	foundUser, err = repo.FindByUsername("john_updated")
	assert.NotNil(t, err, "Expected error when finding deleted user")
	assert.Nil(t, foundUser)
}

func TestUserRepository_FindNonExistentUser(t *testing.T) {
	_, tx := testutils.SetupTestDB()
	defer tx.Rollback()

	repo := user.NewUserRepository(tx)

	// Attempt to find a non-existent user by username
	_, err := repo.FindByUsername("non_existent_user")
	assert.NotNil(t, err)

	// Attempt to find a non-existent user by email
	_, err = repo.FindByEmail("non_existent_user@example.com")
	assert.NotNil(t, err)
}

func TestUserRepository_DeleteNonExistentUser(t *testing.T) {
	_, tx := testutils.SetupTestDB()
	defer tx.Rollback()

	repo := user.NewUserRepository(tx)

	// Attempt to delete a non-existent user
	user := &models.User{
		Username:     "non_existent_user",
		Email:        "non_existent_user@example.com",
		PasswordHash: "hashed_password",
	}

	err := repo.Delete(user)
	assert.NotNil(t, err)
}

func TestUserRepository_HandleEmptyStrings(t *testing.T) {
	_, tx := testutils.SetupTestDB()
	defer tx.Rollback()

	repo := user.NewUserRepository(tx)

	// Attempt to create a user with an empty username
	user := &models.User{
		Username:     "",
		Email:        "empty_username@example.com",
		PasswordHash: "hashed_password",
	}

	err := repo.Create(user)
	assert.NotNil(t, err)

	// Attempt to create a user with an empty email
	user = &models.User{
		Username:     "empty_email",
		Email:        "",
		PasswordHash: "hashed_password",
	}

	err = repo.Create(user)
	assert.NotNil(t, err)
}

func TestUserRepository_UniqueConstraints(t *testing.T) {
	_, tx := testutils.SetupTestDB()
	defer tx.Rollback()

	repo := user.NewUserRepository(tx)

	// Create the first user
	user := &models.User{
		Username:     "unique_user",
		Email:        "unique@example.com",
		PasswordHash: "hashed_password",
	}

	err := repo.Create(user)
	assert.Nil(t, err)

	// Attempt to create another user with the same username
	user2 := &models.User{
		Username:     "unique_user",
		Email:        "another_email@example.com",
		PasswordHash: "hashed_password",
	}

	err = repo.Create(user2)
	assert.NotNil(t, err)

	// Attempt to create another user with the same email
	user3 := &models.User{
		Username:     "another_unique_user",
		Email:        "unique@example.com",
		PasswordHash: "hashed_password",
	}

	err = repo.Create(user3)
	assert.NotNil(t, err)
}
