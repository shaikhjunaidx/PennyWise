package user

import (
	"errors"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/internal/constants"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type UserService struct {
	Repo            UserRepository
	CategoryService UserSignUpCategoryService
	BudgetService   UserSignUpBudgetService
}

func NewUserService(repo UserRepository, categoryService UserSignUpCategoryService, budgetService UserSignUpBudgetService) *UserService {
	return &UserService{
		Repo:            repo,
		CategoryService: categoryService,
		BudgetService:   budgetService,
	}
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

	// category, _ := s.CategoryService.AddCategory(username, constants.DefaultCategoryName, "Default category for uncategorized transactions")

	// s.BudgetService.CreateBudget(username, &category.ID, 0.0, time.Now().Month().String(), time.Now().Year())

	if err := s.addDefaultCategoryAndBudget(user); err != nil {
		return nil, err
	}

	return user, nil
}

// addDefaultCategoryAndBudget handles adding the default category and budget for a new user.
func (s *UserService) addDefaultCategoryAndBudget(user *models.User) error {
	// Add the default category
	category, err := s.CategoryService.AddCategory(user.Username, constants.DefaultCategoryName, "Default category for uncategorized transactions")
	if err != nil {
		return err
	}

	// Add the default budget
	_, err = s.BudgetService.CreateBudget(user.Username, &category.ID, 0.0, time.Now().Format("01"), time.Now().Year())
	if err != nil {
		return err
	}

	return nil
}

// Login authenticates a user based on username and password
func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := ComparePasswords(user.PasswordHash, password); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := GenerateJWTToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
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

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
