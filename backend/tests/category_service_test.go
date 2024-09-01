package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupCategoryService() (*category.CategoryService, *mocks.MockCategoryRepository, *mocks.MockUserRepository) {
	mockCategoryRepo := new(mocks.MockCategoryRepository)
	mockUserRepo := &mocks.MockUserRepository{
		Users: make(map[string]*models.User),
	}

	userService := &user.UserService{Repo: mockUserRepo}

	service := category.NewCategoryService(mockCategoryRepo, userService)
	return service, mockCategoryRepo, mockUserRepo
}

func TestCategoryService_AddCategory(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	category := &models.Category{
		UserID:      user.ID,
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	mockRepo.On("Create", mock.Anything).Return(nil)

	result, err := service.AddCategory(username, category.Name, category.Description)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.UserID, result.UserID)
	assert.Equal(t, category.Name, result.Name)
	assert.Equal(t, category.Description, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetCategoryByID(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	categoryID := uint(1)
	expectedCategory := &models.Category{
		ID:          categoryID,
		UserID:      user.ID,
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	mockRepo.On("FindByID", categoryID).Return(expectedCategory, nil)

	result, err := service.GetCategoryByID(username, categoryID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetAllCategories(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	expectedCategories := []*models.Category{
		{
			UserID:      user.ID,
			Name:        "Groceries",
			Description: "Expenses for groceries",
		},
		{
			UserID:      user.ID,
			Name:        "Utilities",
			Description: "Expenses for utilities",
		},
	}

	mockRepo.On("FindAllByUserID", user.ID).Return(expectedCategories, nil)

	result, err := service.GetAllCategories(username)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, result)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_UpdateCategory(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	categoryID := uint(1)
	existingCategory := &models.Category{
		ID:          categoryID,
		UserID:      user.ID,
		Name:        "Groceries",
		Description: "Expenses for groceries",
	}

	updatedName := "Updated Groceries"
	updatedDescription := "Updated description"

	mockRepo.On("FindByID", categoryID).Return(existingCategory, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)

	result, err := service.UpdateCategory(username, categoryID, updatedName, updatedDescription)

	assert.NoError(t, err)
	assert.Equal(t, updatedName, result.Name)
	assert.Equal(t, updatedDescription, result.Description)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_DeleteCategory(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	user := createTestUser(mockUserRepo, username, 1)

	categoryID := uint(1)
	existingCategory := &models.Category{
		ID:     categoryID,
		UserID: user.ID,
		Name:   "Groceries",
	}

	mockRepo.On("FindByID", categoryID).Return(existingCategory, nil)
	mockRepo.On("DeleteByID", categoryID).Return(nil)

	err := service.DeleteCategory(username, categoryID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCategoryService_FindByName(t *testing.T) {
	service, mockRepo, mockUserRepo := setupCategoryService()

	username := "john_doe"
	createTestUser(mockUserRepo, username, 1)

	categoryName := "Miscellaneous"
	expectedCategory := &models.Category{
		ID:          1,
		Name:        categoryName,
		Description: "Default category for uncategorized transactions",
	}

	mockRepo.On("FindByName", categoryName).Return(expectedCategory, nil)

	result, err := service.Repo.FindByName(categoryName)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)

	mockRepo.AssertExpectations(t)
}
