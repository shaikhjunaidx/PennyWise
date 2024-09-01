package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupCategoryTestRepo(t *testing.T) (*category.CategoryRepositoryImpl, *gorm.DB) {
	_, tx := testutils.SetupTestDB()
	t.Cleanup(func() {
		tx.Rollback()
	})

	return category.NewCategoryRepository(tx), tx
}

func createCategoryRepoTestUser(t *testing.T, tx *gorm.DB, username string) *models.User {
	user := &models.User{
		Username: username,
		Email:    username + "@example.com",
	}
	err := tx.Create(user).Error
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	return user
}

func createTestCategory(t *testing.T, repo *category.CategoryRepositoryImpl, userID uint, name, description string) *models.Category {
	category := &models.Category{
		UserID:      userID,
		Name:        name,
		Description: description,
	}
	err := repo.Create(category)
	assert.NoError(t, err)
	assert.NotZero(t, category.ID)
	return category
}

func TestCategoryRepository_Create(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")
}

func TestCategoryRepository_FindByID(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	category := createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")

	foundCategory, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundCategory)
	assert.Equal(t, category.ID, foundCategory.ID)
}

func TestCategoryRepository_FindAllByUserID(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")
	createTestCategory(t, repo, user.ID, "Utilities", "Expenses for utilities")

	categories, err := repo.FindAllByUserID(user.ID)
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestCategoryRepository_Update(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	category := createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")

	category.Name = "Updated Groceries"
	category.Description = "Updated description"
	err := repo.Update(category)
	assert.NoError(t, err)

	updatedCategory, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Groceries", updatedCategory.Name)
	assert.Equal(t, "Updated description", updatedCategory.Description)
}

func TestCategoryRepository_DeleteByID(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	category := createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")

	err := repo.DeleteByID(category.ID)
	assert.NoError(t, err)

	deletedCategory, err := repo.FindByID(category.ID)
	assert.Error(t, err) // Should return an error since the category should be deleted
	assert.Nil(t, deletedCategory)
}

func TestCategoryRepository_FindByName(t *testing.T) {
	repo, tx := setupCategoryTestRepo(t)

	user := createCategoryRepoTestUser(t, tx, "john_doe")
	createdCategory := createTestCategory(t, repo, user.ID, "Groceries", "Expenses for groceries")

	foundCategory, err := repo.FindByName("Groceries")
	assert.NoError(t, err)
	assert.NotNil(t, foundCategory)
	assert.Equal(t, createdCategory.ID, foundCategory.ID)
	assert.Equal(t, "Groceries", foundCategory.Name)
	assert.Equal(t, "Expenses for groceries", foundCategory.Description)

	nonExistentCategory, err := repo.FindByName("NonExistent")
	assert.Error(t, err)
	assert.Nil(t, nonExistentCategory)
}
