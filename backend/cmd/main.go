package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/db"
	"github.com/shaikhjunaidx/pennywise-backend/internal/routes"
)

func main() {
	database := db.InitDB()
	fmt.Println("Connected to the database:", database.Name())

	router := mux.NewRouter()

	routes.SetupUserRoutes(router, database)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
