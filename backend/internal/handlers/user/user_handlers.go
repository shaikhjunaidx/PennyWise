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
