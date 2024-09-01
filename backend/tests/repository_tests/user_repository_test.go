package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestRepo(t *testing.T) (*user.UserRepositoryImpl, *gorm.DB) {
	_, tx := testutils.SetupTestDB()
	t.Cleanup(func() {
		tx.Rollback()
	})
	return user.NewUserRepository(tx), tx
}

func TestUserRepository_CRUDOperations(t *testing.T) {
	repo, _ := setupTestRepo(t)

	user := &models.User{
		Username:     "john_doe",
		Email:        "john.doe@example.com",
		PasswordHash: "hashed_password",
	}

	assert.NoError(t, repo.Create(user))
	assertUserFoundByUsername(t, repo, "john_doe", user.Email)
	assertUserFoundByEmail(t, repo, "john.doe@example.com", user.Username)

	user.Username = "john_updated"
	assert.NoError(t, repo.Update(user))
	assertUserFoundByUsername(t, repo, "john_updated", user.Email)

	assert.NoError(t, repo.Delete(user))
	assertUserNotFoundByUsername(t, repo, "john_updated")
}

func TestUserRepository_FindNonExistentUser(t *testing.T) {
	repo, _ := setupTestRepo(t)

	assertUserNotFoundByUsername(t, repo, "non_existent_user")
	assertUserNotFoundByEmail(t, repo, "non_existent_user@example.com")
}

func TestUserRepository_DeleteNonExistentUser(t *testing.T) {
	repo, _ := setupTestRepo(t)

	user := &models.User{
		Username:     "non_existent_user",
		Email:        "non_existent_user@example.com",
		PasswordHash: "hashed_password",
	}

	assert.Error(t, repo.Delete(user))
}

func TestUserRepository_HandleEmptyStrings(t *testing.T) {
	repo, _ := setupTestRepo(t)

	user := &models.User{
		Username:     "",
		Email:        "empty_username@example.com",
		PasswordHash: "hashed_password",
	}
	assert.Error(t, repo.Create(user))

	user = &models.User{
		Username:     "empty_email",
		Email:        "",
		PasswordHash: "hashed_password",
	}
	assert.Error(t, repo.Create(user))
}

func TestUserRepository_UniqueConstraints(t *testing.T) {
	repo, _ := setupTestRepo(t)

	user := &models.User{
		Username:     "unique_user",
		Email:        "unique@example.com",
		PasswordHash: "hashed_password",
	}
	assert.NoError(t, repo.Create(user))

	user2 := &models.User{
		Username:     "unique_user",
		Email:        "another_email@example.com",
		PasswordHash: "hashed_password",
	}
	assert.Error(t, repo.Create(user2))

	user3 := &models.User{
		Username:     "another_unique_user",
		Email:        "unique@example.com",
		PasswordHash: "hashed_password",
	}
	assert.Error(t, repo.Create(user3))
}

func assertUserFoundByUsername(t *testing.T, repo *user.UserRepositoryImpl, username, expectedEmail string) {
	foundUser, err := repo.FindByUsername(username)
	assert.NoError(t, err)
	assert.Equal(t, expectedEmail, foundUser.Email)
}

func assertUserFoundByEmail(t *testing.T, repo *user.UserRepositoryImpl, email, expectedUsername string) {
	foundUser, err := repo.FindByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsername, foundUser.Username)
}

func assertUserNotFoundByUsername(t *testing.T, repo *user.UserRepositoryImpl, username string) {
	foundUser, err := repo.FindByUsername(username)
	assert.Error(t, err)
	assert.Nil(t, foundUser)
}

func assertUserNotFoundByEmail(t *testing.T, repo *user.UserRepositoryImpl, email string) {
	foundUser, err := repo.FindByEmail(email)
	assert.Error(t, err)
	assert.Nil(t, foundUser)
}
