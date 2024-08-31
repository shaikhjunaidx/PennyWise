package db

import (
	"fmt"
	"log"
	"os"

	"github.com/shaikhjunaidx/pennywise-backend/internal/config"
	"github.com/shaikhjunaidx/pennywise-backend/internal/constants"
	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	setEnvironment("dev")

	cfg := config.LoadConfig()

	ensureDatabaseExists(cfg)

	db := connectToDatabase(cfg)

	applyMigrations(db)

	if err := EnsureDefaultCategory(db); err != nil {
		log.Fatalf("Failed to ensure default category: %v", err)
	}

	log.Println("Connected to the database and applied migrations")
	return db
}

func setEnvironment(env string) {
	os.Setenv("APP_ENV", env)
}

func ensureDatabaseExists(cfg *config.Config) {
	db := connectToMySQL(cfg)

	if !databaseExists(db, cfg.DBName) {
		if err := createDatabase(db, cfg.DBName); err != nil {
			log.Fatalf("Failed to create database %s: %v", cfg.DBName, err)
		}
		log.Printf("Database %s created", cfg.DBName)
	} else {
		log.Printf("Database %s already exists", cfg.DBName)
	}
}

func connectToMySQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to MySQL server: %v", err)
	}
	return db
}

func databaseExists(db *gorm.DB, dbName string) bool {
	var result int
	db.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&result)
	return result > 0
}

func createDatabase(db *gorm.DB, dbName string) error {
	return db.Exec("CREATE DATABASE " + dbName).Error
}

func connectToDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database %s: %v", cfg.DBName, err)
	}
	return db
}

func applyMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Transaction{}, &models.Budget{}); err != nil {
		log.Fatalf("Could not migrate database schema: %v", err)
	}
}

func EnsureDefaultCategory(db *gorm.DB) error {
	var category models.Category

	err := db.Where("name = ?", constants.DefaultCategoryName).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			defaultCategory := models.Category{
				Name:        constants.DefaultCategoryName,
				Description: "Default category for uncategorized transactions",
			}
			if err := db.Create(&defaultCategory).Error; err != nil {
				return err
			}
			log.Printf("Default category '%s' created successfully.", constants.DefaultCategoryName)
		} else {
			return err
		}
	}
	return nil
}
