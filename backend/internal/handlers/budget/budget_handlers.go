package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/internal/budget"
	"github.com/shaikhjunaidx/pennywise-backend/internal/handlers"
	"github.com/shaikhjunaidx/pennywise-backend/internal/middleware"
	"github.com/shaikhjunaidx/pennywise-backend/models"
)

type BudgetRequest struct {
	CategoryID  *uint   `json:"category_id,omitempty"`
	AmountLimit float64 `json:"amount_limit"`
	BudgetMonth string  `json:"budget_month"`
	BudgetYear  int     `json:"budget_year"`
}

// CreateBudgetHandler handles the creation of a new budget.
// @Summary Create Budget
// @Description Creates a new budget for a user, either overall or for a specific category.
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param   budget  body  handlers.BudgetRequest  true  "Budget Data"
// @Success 201 {object} models.Budget "Created Budget"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets [post]
func CreateBudgetHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Budget

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		createdBudget, err := service.CreateBudget(username, req.CategoryID, req.AmountLimit, req.BudgetMonth, req.BudgetYear)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to create budget", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, createdBudget, http.StatusCreated)
	}
}

// GetBudgetByIDHandler handles retrieving a budget by its ID.
// @Summary Get Budget by ID
// @Description Retrieves a budget by its ID.
// @Tags budgets
// @Produce  json
// @Param   id   path  int  true  "Budget ID"
// @Success 200 {object} models.Budget "Budget"
// @Failure 400 {object} map[string]interface{} "Invalid Budget ID"
// @Failure 404 {object} map[string]interface{} "Budget not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/{id} [get]
func GetBudgetByIDHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		budgetID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid Budget ID", http.StatusBadRequest)
			return
		}

		budget, err := service.GetBudgetByID(uint(budgetID))
		if err != nil {
			handlers.SendErrorResponse(w, "Budget not found", http.StatusNotFound)
			return
		}

		handlers.SendJSONResponse(w, budget, http.StatusOK)
	}
}

// GetBudgetsForUserHandler retrieves all budgets for the authenticated user.
// @Summary Get Budgets for User
// @Description Retrieves all budgets associated with the authenticated user.
// @Tags budgets
// @Produce  json
// @Success 200 {array} models.Budget "List of Budgets"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets [get]
func GetBudgetsForUserHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		budgets, err := service.GetBudgetsForUser(username)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to retrieve budgets", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, budgets, http.StatusOK)
	}
}

type UpdateBudgetRequest struct {
	AmountLimit float64 `json:"amount_limit,omitempty"`
}

// UpdateBudgetHandler handles updating an existing budget.
// @Summary Update Budget
// @Description Updates an existing budget by ID.
// @Tags budgets
// @Accept  json
// @Produce  json
// @Param   id      path  int                    true  "Budget ID"
// @Param   budget  body  handlers.UpdateBudgetRequest  true  "Updated Budget Data"
// @Success 200 {object} models.Budget "Updated Budget"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 404 {object} map[string]interface{} "Budget not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/{id} [put]
func UpdateBudgetHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		budgetID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid Budget ID", http.StatusBadRequest)
			return
		}

		var req struct {
			AmountLimit float64 `json:"amount_limit,omitempty"`
		}

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		budget, err := service.UpdateBudget(uint(budgetID), req.AmountLimit)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to update budget", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, budget, http.StatusOK)
	}
}

// DeleteBudgetHandler handles deleting a budget by its ID.
// @Summary Delete Budget
// @Description Deletes a budget by its ID.
// @Tags budgets
// @Produce  json
// @Param   id   path  int  true  "Budget ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{} "Invalid Budget ID"
// @Failure 404 {object} map[string]interface{} "Budget not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/{id} [delete]
func DeleteBudgetHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		budgetID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid Budget ID", http.StatusBadRequest)
			return
		}

		if err := service.DeleteBudget(uint(budgetID)); err != nil {
			handlers.SendErrorResponse(w, "Failed to delete budget", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetOverallBudgetHandler handles retrieving the overall budget for the authenticated user.
// @Summary Get Overall Budget
// @Description Retrieves the overall budget, including all categories, for the authenticated user.
// @Tags budgets
// @Produce  json
// @Success 200 {object} models.Budget "Overall Budget"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/overall [get]
func GetOverallBudgetHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		overallBudget, err := service.CalculateOverallBudget(username)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to calculate overall budget", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, overallBudget, http.StatusOK)
	}
}

// GetBudgetForUserAndCategoryHandler handles retrieving a budget by category ID for a specific user.
// @Summary Get Budget by Category ID
// @Description Retrieves the budget for the specified category ID for the current month and year for the logged-in user.
// @Tags budgets
// @Produce  json
// @Param categoryID path int true "Category ID"
// @Success 200 {object} models.Budget "Budget"
// @Failure 400 {object} map[string]interface{} "Invalid Category ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Budget not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/category/{categoryID} [get]
func GetBudgetForUserAndCategoryHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		categoryIDStr := vars["category_id"]
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		month := time.Now().Format("January")
		year := time.Now().Year()

		budget, err := service.GetBudgetForUserAndCategory(username, uintPtr(uint(categoryID)), month, year)
		if err != nil {
			handlers.SendErrorResponse(w, "Budget not found", http.StatusNotFound)
			return
		}

		handlers.SendJSONResponse(w, budget, http.StatusOK)
	}
}

// GetBudgetHistoryByCategoryHandler handles retrieving the last 4 months of budget history for a category.
// @Summary Get Budget History by Category
// @Description Retrieves the last 4 months of budget and spending for the given category.
// @Tags budgets
// @Produce  json
// @Param categoryID path int true "Category ID"
// @Success 200 {object} budget.CategoryBudgetHistoryResponse "Budget History"
// @Failure 400 {object} map[string]interface{} "Invalid Category ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/budgets/category/{categoryID}/history [get]
func GetBudgetHistoryByCategoryHandler(service *budget.BudgetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(middleware.UsernameKey).(string)
		if !ok || username == "" {
			handlers.SendErrorResponse(w, "Username not found in context", http.StatusUnauthorized)
			return
		}

		vars := mux.Vars(r)
		categoryIDStr := vars["category_id"]
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}

		budgetHistory, err := service.GetBudgetHistoryForCategory(username, uint(categoryID))
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to retrieve budget history", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, budgetHistory, http.StatusOK)
	}
}

// Helper function to return a pointer to a uint
func uintPtr(i uint) *uint {
	return &i
}
