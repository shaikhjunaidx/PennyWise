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

func createTestCategory(t *testing.T, repo *category.CategoryRepositoryImpl, name, description string) *models.Category {
	category := &models.Category{
		Name:        name,
		Description: description,
	}
	err := repo.Create(category)
	assert.NoError(t, err)
	assert.NotZero(t, category.ID)
	return category
}

func TestCategoryRepository_Create(t *testing.T) {
	repo, _ := setupCategoryTestRepo(t)
	createTestCategory(t, repo, "Groceries", "Expenses for groceries")
}

func TestCategoryRepository_FindByID(t *testing.T) {
	repo, _ := setupCategoryTestRepo(t)

	category := createTestCategory(t, repo, "Groceries", "Expenses for groceries")

	foundCategory, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundCategory)
	assert.Equal(t, category.ID, foundCategory.ID)
}

func TestCategoryRepository_FindAll(t *testing.T) {
	repo, _ := setupCategoryTestRepo(t)

	createTestCategory(t, repo, "Groceries", "Expenses for groceries")
	createTestCategory(t, repo, "Utilities", "Expenses for utilities")

	categories, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, categories, 2)
}

func TestCategoryRepository_Update(t *testing.T) {
	repo, _ := setupCategoryTestRepo(t)

	category := createTestCategory(t, repo, "Groceries", "Expenses for groceries")

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
	repo, _ := setupCategoryTestRepo(t)

	category := createTestCategory(t, repo, "Groceries", "Expenses for groceries")

	err := repo.DeleteByID(category.ID)
	assert.NoError(t, err)

	deletedCategory, err := repo.FindByID(category.ID)
	assert.Error(t, err) // Should return an error since the category should be deleted
	assert.Nil(t, deletedCategory)
}

func TestCategoryRepository_FindByName(t *testing.T) {
    repo, _ := setupCategoryTestRepo(t)

    // Create a test category
    createdCategory := createTestCategory(t, repo, "Groceries", "Expenses for groceries")

    // Attempt to find the category by name
    foundCategory, err := repo.FindByName("Groceries")
    
    // Validate the results
    assert.NoError(t, err)
    assert.NotNil(t, foundCategory)
    assert.Equal(t, createdCategory.ID, foundCategory.ID)
    assert.Equal(t, "Groceries", foundCategory.Name)
    assert.Equal(t, "Expenses for groceries", foundCategory.Description)

    // Attempt to find a category that doesn't exist
    nonExistentCategory, err := repo.FindByName("NonExistent")
    assert.Error(t, err)
    assert.Nil(t, nonExistentCategory)
}