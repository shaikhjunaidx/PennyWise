package test

import (
	"testing"

	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/testutils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupBudgetTestRepo(t *testing.T) (*budget.BudgetRepositoryImpl, *gorm.DB) {
	_, tx := testutils.SetupTestDB()

	t.Cleanup(func() {
		tx.Rollback()
	})

	return budget.NewBudgetRepository(tx), tx
}

func setupTestUserService(db *gorm.DB) *user.UserService {
	userRepo := user.NewUserRepository(db)
	return user.NewUserService(userRepo)
}

func setupTestCategoryService(db *gorm.DB, userService *user.UserService) *category.CategoryService {
	categoryRepo := category.NewCategoryRepository(db)
	return category.NewCategoryService(categoryRepo, userService)
}

// Helper function to create a test budget
func createTestBudget(t *testing.T, repo *budget.BudgetRepositoryImpl, user *models.User, category *models.Category, amountLimit float64) *models.Budget {
	var categoryID *uint
	if category != nil {
		categoryID = &category.ID
	}

	budget := &models.Budget{
		UserID:          user.ID,
		CategoryID:      categoryID,
		AmountLimit:     amountLimit,
		SpentAmount:     0.0,
		RemainingAmount: amountLimit,
		BudgetMonth:     "09",
		BudgetYear:      2024,
	}
	err := repo.Create(budget)
	assert.NoError(t, err)
	assert.NotZero(t, budget.ID)
	return budget
}

func TestBudgetRepository_Create(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and a category
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)

	// Create a budget
	_ = createTestBudget(t, repo, user, category, 1000.0)
}

func TestBudgetRepository_CreateOverallBudget(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)

	// Create a user
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)

	// Create an overall budget (categoryID = nil)
	_ = createTestBudget(t, repo, user, nil, 1000.0)
}

func TestBudgetRepository_FindByID(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and a category
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)

	// Create and find a budget
	createdBudget := createTestBudget(t, repo, user, category, 1000.0)
	foundBudget, err := repo.FindByID(createdBudget.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundBudget)
	assert.Equal(t, createdBudget.ID, foundBudget.ID)
}

func TestBudgetRepository_FindOverallBudget(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)

	// Create a user
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)

	// Create and find an overall budget (categoryID = nil)
	createdBudget := createTestBudget(t, repo, user, nil, 1000.0)
	foundBudget, err := repo.FindByUserIDAndCategoryID(user.ID, nil, "09", 2024)
	assert.NoError(t, err)
	assert.NotNil(t, foundBudget)
	assert.Equal(t, createdBudget.ID, foundBudget.ID)
}

func TestBudgetRepository_FindAllByUserID(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and two categories
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category1, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)
	category2, err := categoryService.AddCategory(user.Username, "Utilities", "Expenses for utilities")
	assert.NoError(t, err)

	// Create multiple budgets
	_ = createTestBudget(t, repo, user, category1, 1000.0)
	_ = createTestBudget(t, repo, user, category2, 1500.0)

	// Find all budgets for the user
	budgets, err := repo.FindAllByUserID(user.ID)
	assert.NoError(t, err)
	assert.Len(t, budgets, 2)
}

func TestBudgetRepository_Update(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and a category
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)

	// Create and update a budget
	budget := createTestBudget(t, repo, user, category, 1000.0)
	budget.AmountLimit = 1200.0
	budget.SpentAmount = 200.0
	budget.RemainingAmount = 1000.0
	err = repo.Update(budget)
	assert.NoError(t, err)

	// Verify the update
	updatedBudget, err := repo.FindByID(budget.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1200.0, updatedBudget.AmountLimit)
	assert.Equal(t, 200.0, updatedBudget.SpentAmount)
	assert.Equal(t, 1000.0, updatedBudget.RemainingAmount)
}

func TestBudgetRepository_DeleteByID(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and a category
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)

	// Create and delete a budget
	budget := createTestBudget(t, repo, user, category, 1000.0)
	err = repo.DeleteByID(budget.ID)
	assert.NoError(t, err)

	// Verify the deletion
	deletedBudget, err := repo.FindByID(budget.ID)
	assert.Error(t, err) // Should return an error since the budget should be deleted
	assert.Nil(t, deletedBudget)
}

func TestBudgetRepository_FindByUserIDAndCategoryID(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)
	categoryService := setupTestCategoryService(db, userService)

	// Create a user and a category
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)
	category, err := categoryService.AddCategory(user.Username, "Groceries", "Expenses for groceries")
	assert.NoError(t, err)

	// Create and find a budget
	createdBudget := createTestBudget(t, repo, user, category, 1000.0)
	foundBudget, err := repo.FindByUserIDAndCategoryID(user.ID, &category.ID, "09", 2024)
	assert.NoError(t, err)
	assert.NotNil(t, foundBudget)
	assert.Equal(t, createdBudget.ID, foundBudget.ID)
}

func TestBudgetRepository_UpdateOverallBudget(t *testing.T) {
	repo, db := setupBudgetTestRepo(t)
	userService := setupTestUserService(db)

	// Create a user
	user, err := userService.SignUp("john_doe", "john.doe@example.com", "password")
	assert.NoError(t, err)

	// Create and update an overall budget (categoryID = nil)
	budget := createTestBudget(t, repo, user, nil, 1000.0)
	budget.AmountLimit = 1200.0
	budget.SpentAmount = 200.0
	budget.RemainingAmount = 1000.0
	err = repo.Update(budget)
	assert.NoError(t, err)

	// Verify the update
	updatedBudget, err := repo.FindByID(budget.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1200.0, updatedBudget.AmountLimit)
	assert.Equal(t, 200.0, updatedBudget.SpentAmount)
	assert.Equal(t, 1000.0, updatedBudget.RemainingAmount)
}