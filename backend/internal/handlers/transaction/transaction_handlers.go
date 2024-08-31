package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/internal/handlers"
	"github.com/shaikhjunaidx/pennywise-backend/internal/transaction"
)

// CreateTransactionHandler handles the creation of a new transaction.
// @Summary Create Transaction
// @Description Creates a new transaction for the authenticated user, linking it to a specific category.
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   user_id         body  uint    true  "User ID"
// @Param   category_id     body  uint    true  "Category ID"
// @Param   amount          body  float64 true  "Amount"
// @Param   description     body  string  false "Description"
// @Param   transaction_date body string true  "Transaction Date in RFC3339 format"
// @Success 201 {object} models.Transaction "Created Transaction"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions [post]
func CreateTransactionHandler(service *transaction.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID          uint    `json:"user_id"`
			CategoryID      uint    `json:"category_id"`
			Amount          float64 `json:"amount"`
			Description     string  `json:"description"`
			TransactionDate string  `json:"transaction_date"`
		}

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		transaction, err := service.AddTransaction(req.UserID, req.CategoryID, req.Amount, req.Description, transactionDate)
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to create transaction", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, transaction, http.StatusCreated)

	}
}

// GetTransactionsHandler handles retrieving transactions for the authenticated user.
// @Summary Get Transactions
// @Description Retrieves a list of transactions for the authenticated user, with optional filtering by date, category, etc.
// @Tags transactions
// @Produce  json
// @Param   user_id  query uint true "User ID"
// @Success 200 {array} models.Transaction "List of Transactions"
// @Failure 400 {object} map[string]interface{} "Invalid or missing user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions [get]
func GetTransactionsHandler(service *transaction.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("user_id")
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil || userID == 0 {
			handlers.SendErrorResponse(w, "Invalid or missing user ID", http.StatusBadRequest)
			return
		}

		transactions, err := service.GetTransactionsForUser(uint(userID))
		if err != nil {
			handlers.SendErrorResponse(w, "Failed to retrieve transactions", http.StatusInternalServerError)
			return
		}

		handlers.SendJSONResponse(w, transactions, http.StatusOK)
	}
}

// GetTransactionByIDHandler handles retrieving a single transaction by its ID.
// @Summary Get Transaction by ID
// @Description Retrieves a single transaction by its ID for the authenticated user.
// @Tags transactions
// @Produce  json
// @Param   id  path uint true "Transaction ID"
// @Success 200 {object} models.Transaction "Transaction data"
// @Failure 400 {object} map[string]interface{} "Invalid transaction ID"
// @Failure 404 {object} map[string]interface{} "Transaction not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions/{id} [get]
func GetTransactionByIDHandler(service *transaction.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transactionID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil || transactionID == 0 {
			handlers.SendErrorResponse(w, "Invalid transaction ID", http.StatusBadRequest)
			return
		}

		transaction, err := service.GetTransactionByID(uint(transactionID))
		if err != nil {
			if err.Error() == "record not found" {
				handlers.SendErrorResponse(w, "Transaction not found", http.StatusNotFound)
			} else {
				handlers.SendErrorResponse(w, "Failed to retrieve transaction", http.StatusInternalServerError)
			}
			return
		}

		handlers.SendJSONResponse(w, transaction, http.StatusOK)
	}
}

// UpdateTransactionHandler handles updating an existing transaction.
// @Summary Update Transaction
// @Description Updates an existing transaction, allowing changes to the amount, category, description, or date.
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   id            path  uint    true  "Transaction ID"
// @Param   amount        body  float64 true  "Amount"
// @Param   category_id   body  uint    true  "Category ID"
// @Param   description   body  string  false "Description"
// @Param   transaction_date body string true "Transaction Date in RFC3339 format"
// @Success 200 {object} models.Transaction "Updated Transaction"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 404 {object} map[string]interface{} "Transaction not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions/{id} [put]
func UpdateTransactionHandler(service *transaction.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Amount          float64 `json:"amount"`
			CategoryID      uint    `json:"category_id"`
			Description     string  `json:"description"`
			TransactionDate string  `json:"transaction_date"`
		}

		vars := mux.Vars(r)
		transactionID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid transaction ID", http.StatusBadRequest)
			return
		}

		if err := handlers.ParseJSONRequest(w, r, &req); err != nil {
			return
		}

		transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		transaction, err := service.UpdateTransaction(uint(transactionID), req.Amount, req.CategoryID, req.Description, transactionDate)
		if err != nil {
			if err.Error() == "record not found" {
				handlers.SendErrorResponse(w, "Transaction not found", http.StatusNotFound)
			} else {
				handlers.SendErrorResponse(w, "Failed to update transaction", http.StatusInternalServerError)
			}
			return
		}

		handlers.SendJSONResponse(w, transaction, http.StatusOK)
	}
}

// DeleteTransactionHandler handles deleting an existing transaction.
// @Summary Delete Transaction
// @Description Deletes a specific transaction by its ID.
// @Tags transactions
// @Param   id  path  uint  true  "Transaction ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{} "Invalid transaction ID"
// @Failure 404 {object} map[string]interface{} "Transaction not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/transactions/{id} [delete]
func DeleteTransactionHandler(service *transaction.TransactionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transactionID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			handlers.SendErrorResponse(w, "Invalid transaction ID", http.StatusBadRequest)
			return
		}

		if err := service.DeleteTransaction(uint(transactionID)); err != nil {
			if err.Error() == "record not found" {
				handlers.SendErrorResponse(w, "Transaction not found", http.StatusNotFound)
			} else {
				handlers.SendErrorResponse(w, "Failed to delete transaction", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
