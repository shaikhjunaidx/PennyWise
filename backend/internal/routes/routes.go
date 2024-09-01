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

func SetupUserRoutes(router *mux.Router, db *gorm.DB) {
	userRepo := user.NewUserRepository(db)
	userService := &user.UserService{Repo: userRepo}

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api/signup", userHandlers.SignUpHandler(userService)).Methods("POST")
	router.HandleFunc("/api/login", userHandlers.LoginHandler(userService)).Methods("POST")
}

func SetupTransactionRoutes(router *mux.Router, db *gorm.DB) {
	userRepo := user.NewUserRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)
	budgetRepo := budget.NewBudgetRepository(db)
	userService := user.NewUserService(userRepo)
	budgetService := budget.NewBudgetService(budgetRepo, userService)

	transactionService := transaction.NewTransactionService(transactionRepo, userRepo, categoryRepo, budgetService)

	transactionRouter := router.PathPrefix("/api/transactions").Subrouter()
	transactionRouter.Use(middleware.JWTMiddleware)

	transactionRouter.HandleFunc("", transactionHandlers.CreateTransactionHandler(transactionService)).Methods("POST")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.UpdateTransactionHandler(transactionService)).Methods("PUT")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.DeleteTransactionHandler(transactionService)).Methods("DELETE")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.GetTransactionByIDHandler(transactionService)).Methods("GET")
	transactionRouter.HandleFunc("", transactionHandlers.GetTransactionsHandler(transactionService)).Methods("GET")
}

func SetupCategoryRoutes(router *mux.Router, db *gorm.DB) {
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)

	categoryRouter := router.PathPrefix("/api/categories").Subrouter()

	categoryRouter.HandleFunc("", categoryHandlers.CreateCategoryHandler(categoryService)).Methods("POST")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.GetCategoryByIDHandler(categoryService)).Methods("GET")
	categoryRouter.HandleFunc("", categoryHandlers.GetAllCategoriesHandler(categoryService)).Methods("GET")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.UpdateCategoryHandler(categoryService)).Methods("PUT")
	categoryRouter.HandleFunc("/{id:[0-9]+}", categoryHandlers.DeleteCategoryHandler(categoryService)).Methods("DELETE")
}

func SetupBudgetRoutes(router *mux.Router, db *gorm.DB) {
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	budgetRepo := budget.NewBudgetRepository(db)
	budgetService := budget.NewBudgetService(budgetRepo, userService)

	budgetRouter := router.PathPrefix("/api/budgets").Subrouter()
	budgetRouter.Use(middleware.JWTMiddleware)

	budgetRouter.HandleFunc("", budgetHandlers.CreateBudgetHandler(budgetService)).Methods("POST")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.GetBudgetByIDHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("", budgetHandlers.GetBudgetsForUserHandler(budgetService)).Methods("GET")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.UpdateBudgetHandler(budgetService)).Methods("PUT")
	budgetRouter.HandleFunc("/{id:[0-9]+}", budgetHandlers.DeleteBudgetHandler(budgetService)).Methods("DELETE")
	budgetRouter.HandleFunc("/overall", budgetHandlers.GetOverallBudgetHandler(budgetService)).Methods("GET")
}
