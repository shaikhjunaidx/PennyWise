package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"github.com/shaikhjunaidx/pennywise-backend/utils"
)

// SeedTestDataAndVerify seeds test data and validates using the API.
func SeedTestDataAndVerify() {
	env := os.Getenv("ENV")
	if env != "development" {
		return
	}

	log.Println("Seeding test data for development environment...")

	// Test User
	testUser := map[string]string{
		"username": "test",
		"email":    "test@example.com",
		"password": "123",
	}

	// Create User via API
	resp, err := utils.MakeAPICall("POST", "http://localhost:8080/api/signup", "", testUser)
	if err != nil {
		log.Fatalf("Error creating test user: %v", err)
	}

	testUserLogin := map[string]string{
		"username": "test",
		"password": "123",
	}

	// Login User via API
	// Create User via API
	resp, err = utils.MakeAPICall("POST", "http://localhost:8080/api/login", "", testUserLogin)
	if err != nil {
		log.Fatalf("Error logging in  test user: %v", err)
	}

	// Extract JWT token from the response
	var signupResponse map[string]string
	err = json.Unmarshal(resp, &signupResponse)
	if err != nil || signupResponse["token"] == "" {
		log.Fatalf("Failed to get JWT token: %v", err)
	}
	token := signupResponse["token"]

	// Create Categories via API using the token
	categories := []map[string]string{
		{"name": "Groceries", "description": "Test Groceries category"},
		{"name": "Utilities", "description": "Test Utilities category"},
	}

	for _, category := range categories {
		_, err := utils.MakeAPICall("POST", "http://localhost:8080/api/categories", token, category)
		if err != nil {
			log.Fatalf("Error creating category: %v", err)
		}
	}

	// Create Budgets via API using the token
	budgets := []map[string]interface{}{
		{"category_id": 2, "amount_limit": 1000.0, "budget_month": "09", "budget_year": 2024},
		{"category_id": 2, "amount_limit": 1000.0, "budget_month": "08", "budget_year": 2024},
		{"category_id": 3, "amount_limit": 1500.0, "budget_month": "08", "budget_year": 2024},
		{"category_id": 3, "amount_limit": 1500.0, "budget_month": "09", "budget_year": 2024},
	}

	for _, budget := range budgets {
		_, err := utils.MakeAPICall("POST", "http://localhost:8080/api/budgets", token, budget)
		if err != nil {
			log.Fatalf("Error creating budget: %v", err)
		}
	}

	// Verify Budgets via API using the token
	resp, err = utils.MakeAPICall("GET", "http://localhost:8080/api/budgets", token, nil)
	if err != nil {
		log.Fatalf("Error fetching budgets: %v", err)
	}

	var fetchedBudgets []models.Budget
	if err := json.Unmarshal(resp, &fetchedBudgets); err != nil {
		log.Fatalf("Error decoding budgets: %v", err)
	}

	// fmt.Printf("Budgets verified: %+v\n", fetchedBudgets)

	// Create Transactions via API using the token
	transactions := []map[string]interface{}{
		{"category_id": 2, "amount": 100.0, "description": "Groceries Transaction", "transaction_date": "2024-09-08T00:00:00Z"},
		{"category_id": 2, "amount": 200.0, "description": "Groceries Transaction", "transaction_date": "2024-09-01T00:00:00Z"},
		{"category_id": 2, "amount": 350.0, "description": "Groceries Transaction", "transaction_date": "2024-08-25T00:00:00Z"},
		{"category_id": 2, "amount": 150.0, "description": "Groceries Transaction", "transaction_date": "2024-08-18T00:00:00Z"},
		{"category_id": 3, "amount": 200.0, "description": "Utilities Transaction", "transaction_date": "2024-09-08T00:00:00Z"},
		{"category_id": 3, "amount": 350.0, "description": "Groceries Transaction", "transaction_date": "2024-08-25T00:00:00Z"},
		{"category_id": 3, "amount": 150.0, "description": "Groceries Transaction", "transaction_date": "2024-08-18T00:00:00Z"},
		{"category_id": 3, "amount": 200.0, "description": "Groceries Transaction", "transaction_date": "2024-09-01T00:00:00Z"},
	}

	for _, transaction := range transactions {
		_, err := utils.MakeAPICall("POST", "http://localhost:8080/api/transactions", token, transaction)
		if err != nil {
			log.Fatalf("Error creating transaction: %v", err)
		}
	}

	// Verify Transactions via API using the token
	resp, err = utils.MakeAPICall("GET", "http://localhost:8080/api/transactions", token, nil)
	if err != nil {
		log.Fatalf("Error fetching transactions: %v", err)
	}

	var fetchedTransactions []models.Transaction
	if err := json.Unmarshal(resp, &fetchedTransactions); err != nil {
		log.Fatalf("Error decoding transactions: %v", err)
	}

	// fmt.Printf("Transactions verified: %+v\n", fetchedTransactions)

	log.Println("Test data and API verification completed.")
}
