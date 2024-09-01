package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/internal/category"
	"github.com/shaikhjunaidx/pennywise-backend/internal/handlers"
	"github.com/shaikhjunaidx/pennywise-backend/internal/middleware"
)

// CategoryRequest struct is used for decoding JSON requests
type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateCategoryHandler handles the creation of a new category.
// @Summary Create Category
// @Description Creates a new category for transactions and budgets.
// @Tags categories
// @Accept  json
// @Produce  json
// @Param   category  body  handlers.CategoryRequest  true  "Category"
// @Success 201 {object} models.Category "Created Category"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories [post]
func CreateCategoryHandler(service *category.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		var req CategoryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			handlers.SendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		category, err := service.AddCategory(username, req.Name, req.Description)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to create category", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, category, http.StatusCreated)
	}
}

// GetCategoryByIDHandler handles retrieving a category by its ID.
// @Summary Get Category by ID
// @Description Retrieves a category by its ID.
// @Tags categories
// @Produce  json
// @Param   id   path  int  true  "Category ID"
// @Success 200 {object} models.Category "Category"
// @Failure 400 {object} map[string]interface{} "Invalid Category ID"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id} [get]
func GetCategoryByIDHandler(service *category.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil || id == 0 {
			handlers.SendErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		category, err := service.GetCategoryByID(username, uint(id))
		if err != nil {
			handlers.SendErrorResponse(w, "Category not found", http.StatusNotFound)
			return
		}

		handlers.SendJSONResponse(w, category, http.StatusOK)
	}
}

// GetAllCategoriesHandler handles retrieving all categories for the user.
// @Summary Get All Categories
// @Description Retrieves all categories for the authenticated user.
// @Tags categories
// @Produce  json
// @Success 200 {array} models.Category "List of Categories"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories [get]
func GetAllCategoriesHandler(service *category.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		categories, err := service.GetAllCategories(username)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to retrieve categories", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, categories, http.StatusOK)
	}
}

// UpdateCategoryHandler handles updating an existing category.
// @Summary Update Category
// @Description Updates an existing category by ID.
// @Tags categories
// @Accept  json
// @Produce  json
// @Param   id          path  int                      true  "Category ID"
// @Param   category    body  handlers.CategoryRequest false "Category"
// @Success 200 {object} models.Category "Updated Category"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id} [put]
func UpdateCategoryHandler(service *category.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil || id == 0 {
			handlers.SendErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		var req CategoryRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			handlers.SendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		category, err := service.UpdateCategory(username, uint(id), req.Name, req.Description)
		if err != nil {
			handlers.SendErrorResponse(w, "Category not found", http.StatusNotFound)
			return
		}

		handlers.SendJSONResponse(w, category, http.StatusOK)
	}
}

// DeleteCategoryHandler handles deleting a category by its ID.
// @Summary Delete Category
// @Description Deletes a category by its ID.
// @Tags categories
// @Produce  json
// @Param   id   path  int  true  "Category ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{} "Invalid Category ID"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/categories/{id} [delete]
func DeleteCategoryHandler(service *category.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil || id == 0 {
			handlers.SendErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		if err := service.DeleteCategory(username, uint(id)); err != nil {
			handlers.SendErrorResponse(w, "Category not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
