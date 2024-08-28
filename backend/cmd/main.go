package main

import (
	"fmt"

	"github.com/shaikhjunaidx/pennywise-backend/db"
)

func main() {
	database := db.InitDB()
	fmt.Println(database.Name())
}
