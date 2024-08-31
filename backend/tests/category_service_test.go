package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func setupCategoryService() (*category.CategoryService, *mocks.MockCategoryRepository) {
	mockRepo := new(mocks.MockCategoryRepository)
	service := category.NewCategoryService(mockRepo)
	return service, mockRepo
}

func TestCategoryService_AddCategory(t *testing.T) {
	service, mockRepo := setupCategoryService()

	category := &models.Category{
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	mockRepo.On("Create", category).Return(nil)

	result, err := service.AddCategory(category.Name, category.Description)

	assert.NoError(t, err)
	assert.Equal(t, category, result)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetCategoryByID(t *testing.T) {
	service, mockRepo := setupCategoryService()

	categoryID := uint(1)
	expectedCategory := &models.Category{
		ID:          categoryID,
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	mockRepo.On("FindByID", categoryID).Return(expectedCategory, nil)

	result, err := service.GetCategoryByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetAllCategories(t *testing.T) {
	service, mockRepo := setupCategoryService()

	expectedCategories := []*models.Category{
		{
			Name:        "Groceries",
			Description: "Expenses for groceries",
		},
		{
			Name:        "Utilities",
			Description: "Expenses for utilities",
		},
	}

	mockRepo.On("FindAll").Return(expectedCategories, nil)

	result, err := service.GetAllCategories()

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, result)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_UpdateCategory(t *testing.T) {
	service, mockRepo := setupCategoryService()

	categoryID := uint(1)
	existingCategory := &models.Category{
		ID:          categoryID,
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	updatedName := "Updated Groceries"
	updatedDescription := "Updated description"

	mockRepo.On("FindByID", categoryID).Return(existingCategory, nil)
	mockRepo.On("Update", existingCategory).Return(nil)

	result, err := service.UpdateCategory(categoryID, updatedName, updatedDescription)

	assert.NoError(t, err)
	assert.Equal(t, updatedName, result.Name)
	assert.Equal(t, updatedDescription, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_DeleteCategory(t *testing.T) {
	service, mockRepo := setupCategoryService()

	categoryID := uint(1)

	mockRepo.On("DeleteByID", categoryID).Return(nil)

	err := service.DeleteCategory(categoryID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
