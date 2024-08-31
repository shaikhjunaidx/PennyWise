package handlers

import (
	"net/http"

	"github.com/shaikhjunaidx/pennywise-backend/internal/handlers"
	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"password123"`
}

type SignUpRequest struct {
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"password123"`
}

// SignUpHandler handles user registration requests.
// @Summary User Registration
// @Description Registers a new user with the given username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   signupData  body  SignUpRequest  true  "Sign Up Data"
// @Success 201 {object} UserResponse "Created User"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/signup [post]
func SignUpHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		user, err := s.SignUp(req.Username, req.Email, req.Password)
		if err != nil {
			handlers.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, user, http.StatusCreated)
	}
}

// LoginHandler handles user login requests.
// @Summary User Login
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   loginData  body  LoginRequest  true  "Login Data"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/login [post]
func LoginHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		token, err := s.Login(req.Username, req.Password)
		if err != nil {
			handlers.SendErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		handlers.SendJSONResponse(w, map[string]string{"token": token}, http.StatusOK)
	}
}
