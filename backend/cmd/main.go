package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/db"
	_ "github.com/shaikhjunaidx/pennywise-backend/docs"
	"github.com/shaikhjunaidx/pennywise-backend/internal/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	database := db.InitDB()
	fmt.Println("Connected to the database:", database.Name())

	router := mux.NewRouter()

	// Set up routes for the application
	routes.SetupUserRoutes(router, database)

	// Serve the Swagger UI at /swagger/
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Serve the Swagger JSON documentation file
	router.Handle("/swagger/doc.json", http.FileServer(http.Dir("./docs")))

	// Set up CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
}
