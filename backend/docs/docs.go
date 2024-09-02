// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/budgets": {
            "get": {
                "description": "Retrieves all budgets associated with the authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Get Budgets for User",
                "responses": {
                    "200": {
                        "description": "List of Budgets",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Budget"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new budget for a user, either overall or for a specific category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Create Budget",
                "parameters": [
                    {
                        "description": "Budget Data",
                        "name": "budget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created Budget",
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/budgets/category/{categoryID}": {
            "get": {
                "description": "Retrieves the budget for the specified category ID for the current month and year for the logged-in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Get Budget by Category ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "categoryID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Budget",
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    },
                    "400": {
                        "description": "Invalid Category ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Budget not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/budgets/overall": {
            "get": {
                "description": "Retrieves the overall budget, including all categories, for the authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Get Overall Budget",
                "responses": {
                    "200": {
                        "description": "Overall Budget",
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/budgets/{id}": {
            "get": {
                "description": "Retrieves a budget by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Get Budget by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Budget",
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    },
                    "400": {
                        "description": "Invalid Budget ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Budget not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing budget by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Update Budget",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Budget Data",
                        "name": "budget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated Budget",
                        "schema": {
                            "$ref": "#/definitions/models.Budget"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Budget not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a budget by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "budgets"
                ],
                "summary": "Delete Budget",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Budget ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid Budget ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Budget not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/categories": {
            "get": {
                "description": "Retrieves all categories for the authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get All Categories",
                "responses": {
                    "200": {
                        "description": "List of Categories",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new category for transactions and budgets.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Create Category",
                "parameters": [
                    {
                        "description": "Category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created Category",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/categories/{id}": {
            "get": {
                "description": "Retrieves a category by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get Category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid Category ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Category not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing category by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Update Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category",
                        "name": "category",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handlers.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated Category",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Category not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a category by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Delete Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid Category ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Category not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Data",
                        "name": "loginData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/signup": {
            "post": {
                "description": "Registers a new user with the given username, email, and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "Sign Up Data",
                        "name": "signupData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created User",
                        "schema": {
                            "$ref": "#/definitions/handlers.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/transactions": {
            "get": {
                "description": "Retrieves a list of transactions for the authenticated user, with optional filtering by date, category, etc.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get Transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of Transactions",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid or missing user ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new transaction for the authenticated user, linking it to a specific category.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Create Transaction",
                "parameters": [
                    {
                        "description": "Transaction Data",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created Transaction",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/transactions/{id}": {
            "get": {
                "description": "Retrieves a single transaction by its ID for the authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get Transaction by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transaction data",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Invalid transaction ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Transaction not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing transaction, allowing changes to the amount, category, description, or date.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Update Transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Transaction Data",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated Transaction",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Transaction not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a specific transaction by its ID.",
                "tags": [
                    "transactions"
                ],
                "summary": "Delete Transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid transaction ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Transaction not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CategoryRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },
        "handlers.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "john_doe"
                }
            }
        },
        "handlers.TransactionRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "transaction_date": {
                    "type": "string"
                }
            }
        },
        "handlers.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Budget": {
            "type": "object",
            "properties": {
                "amount_limit": {
                    "type": "number"
                },
                "budget_month": {
                    "type": "string"
                },
                "budget_year": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "remaining_amount": {
                    "type": "number"
                },
                "spent_amount": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "transaction_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
