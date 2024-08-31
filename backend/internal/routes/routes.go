package routes

import (
	"github.com/gorilla/mux"

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

func SetupTransactionROutes(router *mux.Router, db *gorm.DB) {
	transactionRepo := transaction.NewTransactionRepository(db)
	transactionService := transaction.NewTransactionService(transactionRepo)

	transactionRouter := router.PathPrefix("/api/transactions").Subrouter()
	transactionRouter.Use(middleware.JWTMiddleware)

	transactionRouter.HandleFunc("", transactionHandlers.CreateTransactionHandler(transactionService)).Methods("POST")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.UpdateTransactionHandler(transactionService)).Methods("PUT")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.DeleteTransactionHandler(transactionService)).Methods("DELETE")
	transactionRouter.HandleFunc("/{id:[0-9]+}", transactionHandlers.GetTransactionByIDHandler(transactionService)).Methods("GET")
	transactionRouter.HandleFunc("", transactionHandlers.GetTransactionsHandler(transactionService)).Methods("GET")

}
