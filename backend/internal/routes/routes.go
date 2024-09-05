package routes

import (
	"github.com/gorilla/mux"

	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	budgetHandlers "github.com/shaikhjunaidx/pennywise-backend/internal/handlers/budget"
	categoryHandlers "github.com/shaikhjunaidx/pennywise-backend/internal/handlers/category"
	transactionHandlers "github.com/shaikhjunaidx/pennywise-backend/internal/handlers/transaction"
	userHandlers "github.com/shaikhjunaidx/pennywise-backend/internal/handlers/user"
	"github.com/shaikhjunaidx/pennywise-backend/internal/middleware"
	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// Initialize repositories
func initRepositories(db *gorm.DB) (user.UserRepository, category.CategoryRepository, budget.BudgetRepository, transaction.TransactionRepository) {
	return user.NewUserRepository(db), category.NewCategoryRepository(db), budget.NewBudgetRepository(db), transaction.NewTransactionRepository(db)
}

// Initialize services
func initServices(db *gorm.DB) (*user.UserService, *category.CategoryService, *budget.BudgetService, *transaction.TransactionService) {
	userRepo, categoryRepo, budgetRepo, transactionRepo := initRepositories(db)

	userService := &user.UserService{Repo: userRepo}
	categoryService := category.NewCategoryService(categoryRepo, userService)
	budgetService := budget.NewBudgetService(budgetRepo, userService)
	transactionService := transaction.NewTransactionService(transactionRepo, userRepo, categoryRepo, budgetService)

	userService.CategoryService = categoryService
	userService.BudgetService = budgetService

	return userService, categoryService, budgetService, transactionService
}

func SetupUserRoutes(router *mux.Router, db *gorm.DB) {
	userService, _, _, _ := initServices(db)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api/signup", userHandlers.SignUpHandler(userService)).Methods("POST")
	router.HandleFunc("/api/login", userHandlers.LoginHandler(userService)).Methods("POST")
}

func SetupTransactionRoutes(router *mux.Router, db *gorm.DB) {
	_, _, _, transactionService := initServices(db)

	transactionRouter := router.PathPrefix("/api/transactions").Subrouter()
	transactionRouter.Use(middleware.JWTMiddleware)

	transactionRouter.HandleFunc("", transactionHandlers.CreateTransactionHandler(transactionService)).Methods("POST")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.UpdateTransactionHandler(transactionService)).Methods("PUT")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.DeleteTransactionHandler(transactionService)).Methods("DELETE")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.GetTransactionByIDHandler(transactionService)).Methods("GET")
	transactionRouter.HandleFunc("", transactionHandlers.GetTransactionsHandler(transactionService)).Methods("GET")
	transactionRouter.HandleFunc("/category/{category_id:[0-9]+}", transactionHandlers.GetTransactionsByCategoryHandler(transactionService)).Methods("GET")
	transactionRouter.HandleFunc("/weekly", transactionHandlers.GetWeeklySpendingHandler(transactionService)).Methods("GET")

}

func SetupCategoryRoutes(router *mux.Router, db *gorm.DB) {
	_, categoryService, _, _ := initServices(db)

	categoryRouter := router.PathPrefix("/api/categories").Subrouter()
	categoryRouter.Use(middleware.JWTMiddleware)

	categoryRouter.HandleFunc("", categoryHandlers.CreateCategoryHandler(categoryService)).Methods("POST")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.GetCategoryByIDHandler(categoryService)).Methods("GET")
	categoryRouter.HandleFunc("", categoryHandlers.GetAllCategoriesHandler(categoryService)).Methods("GET")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.UpdateCategoryHandler(categoryService)).Methods("PUT")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.DeleteCategoryHandler(categoryService)).Methods("DELETE")
}

func SetupBudgetRoutes(router *mux.Router, db *gorm.DB) {
	_, _, budgetService, _ := initServices(db)

	budgetRouter := router.PathPrefix("/api/budgets").Subrouter()
	budgetRouter.Use(middleware.JWTMiddleware)

	budgetRouter.HandleFunc("", budgetHandlers.CreateBudgetHandler(budgetService)).Methods("POST")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.GetBudgetByIDHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("", budgetHandlers.GetBudgetsForUserHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.UpdateBudgetHandler(budgetService)).Methods("PUT")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.DeleteBudgetHandler(budgetService)).Methods("DELETE")
	budgetRouter.HandleFunc("/overall", budgetHandlers.GetOverallBudgetHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("/category/{category_id:[0-9]+}", budgetHandlers.GetBudgetForUserAndCategoryHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("/category/{category_id:[0-9]+}/history", budgetHandlers.GetBudgetHistoryByCategoryHandler(budgetService)).Methods("GET")
}
