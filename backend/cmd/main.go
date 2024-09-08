package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/db"
	_ "github.com/shaikhjunaidx/pennywise-backend/docs"
	"github.com/shaikhjunaidx/pennywise-backend/internal/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	os.Setenv("ENV", "development")

	database := db.InitDB()

	fmt.Println("Connected to the database:", database.Name())

	router := mux.NewRouter()

	routes.SetupUserRoutes(router, database)
	routes.SetupTransactionRoutes(router, database)
	routes.SetupCategoryRoutes(router, database)
	routes.SetupBudgetRoutes(router, database)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/swagger/doc.json", http.FileServer(http.Dir("./docs")))

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	go func() {
		log.Println("Server is running on port 8080...")
		if err := http.ListenAndServe(":8080", corsMiddleware(router)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Seed the test data and make API validation calls after the server is up
	db.SeedTestDataAndVerify()

	select {}
}
