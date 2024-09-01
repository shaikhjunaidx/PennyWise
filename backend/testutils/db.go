package testutils

import (
	"fmt"
	"log"
	"os"

	"github.com/shaikhjunaidx/pennywise-backend/internal/config"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, *gorm.DB) {
	setTestEnvironment()

	cfg := config.LoadConfig()

	// Connect to MySQL server and ensure the test database exists
	db := connectToMySQL(cfg)
	ensureDatabaseExists(db, cfg.DBName)

	// Connect to the test database and apply migrations
	testDB := connectToTestDB(cfg)
	applyMigrations(testDB)

	// Begin a transaction
	tx := testDB.Begin()

	return testDB, tx
}

func setTestEnvironment() {
	os.Setenv("APP_ENV", "test")
}

func connectToMySQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}
	return db
}

func ensureDatabaseExists(db *gorm.DB, dbName string) {
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error; err != nil {
		log.Fatalf("Failed to create or verify database %s: %v", dbName, err)
	}
}

func connectToTestDB(cfg *config.Config) *gorm.DB {
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the test database %s: %v", cfg.DBName, err)
	}
	return db
}

func applyMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Transaction{}, &models.Budget{}); err != nil {
		log.Fatalf("Could not migrate database schema: %v", err)
	}
}
