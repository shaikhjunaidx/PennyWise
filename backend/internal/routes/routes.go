package routes

import (
	"github.com/gorilla/mux"

	handlers "github.com/shaikhjunaidx/pennywise-backend/internal/handlers/user"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *mux.Router, db *gorm.DB) {
	userRepo := user.NewUserRepository(db)
	userService := &user.UserService{Repo: userRepo}

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api/signup", handlers.SignUpHandler(userService)).Methods("POST")
	router.HandleFunc("/api/login", handlers.LoginHandler(userService)).Methods("POST")
}
