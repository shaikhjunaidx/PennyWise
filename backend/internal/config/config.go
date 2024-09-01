package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() *Config {
	env := getEnvironment()
	envFile := getEnvFilePath(env)

	fmt.Printf("Loading environment: %v \n", env)

	loadEnvFile(envFile)

	return &Config{
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
	}
}

func getEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		log.Fatal("APP_ENV is not set")
	}
	return env
}

func getEnvFilePath(env string) string {
	if env == "test" {
		return fmt.Sprintf("../../.env.%s", env)
	}
	return fmt.Sprintf(".env.%s", env)
}

func loadEnvFile(envFile string) {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading .env file for environment: %v", envFile)
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
