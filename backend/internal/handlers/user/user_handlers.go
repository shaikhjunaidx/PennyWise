package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shaikhjunaidx/pennywise-backend/internal/user"
)

func parseJSONRequest(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}
	return nil
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// SignUpHandler handles user registration requests.
// @Summary User Registration
// @Description Registers a new user with the given username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Param   email     body  string  true  "Email"
// @Param   password  body  string  true  "Password"
// @Success 201 {object} UserResponse "Created User"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/signup [post]
func SignUpHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := parseJSONRequest(w, r, &req); err != nil {
			return
		}

		user, err := s.SignUp(req.Username, req.Email, req.Password)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sendJSONResponse(w, user, http.StatusCreated)
	}
}

// LoginHandler handles user login requests.
// @Summary User Login
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   username  body  string  true  "Username"
// @Param   password  body  string  true  "Password"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /api/login [post]
func LoginHandler(s *user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := parseJSONRequest(w, r, &req); err != nil {
			return
		}

		token, err := s.Login(req.Username, req.Password)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sendJSONResponse(w, map[string]string{"token": token}, http.StatusOK)
	}
}
