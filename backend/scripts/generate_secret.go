package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

func generateSecretKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func ensureEnvFileExists(envFilePath string) error {
	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		file, err := os.Create(envFilePath)
		if err != nil {
			return fmt.Errorf("failed to create .env file: %v", err)
		}
		defer file.Close()
		fmt.Println(".env file created")
	}
	return nil
}

func writeSecretToEnvFile(secretKey, envFilePath string) error {
	data, err := os.ReadFile(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read .env file: %v", err)
	}

	if strings.Contains(string(data), "JWT_SECRET=") {
		return fmt.Errorf("JWT_SECRET already exists in .env file")
	}

	// Append the JWT_SECRET to the file
	file, err := os.OpenFile(envFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .env file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("\nJWT_SECRET=%s\n", secretKey)); err != nil {
		return fmt.Errorf("failed to write secret to .env file: %v", err)
	}

	return nil
}

func main() {
	// Use a command-line flag to specify the environment file
	envFilePath := flag.String("env", ".env", "Path to the environment file (e.g., .env.dev, .env.test)")
	flag.Parse()

	secretKey, err := generateSecretKey(32)
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}

	if err := ensureEnvFileExists(*envFilePath); err != nil {
		fmt.Println(err)
		return
	}

	if err := writeSecretToEnvFile(secretKey, *envFilePath); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("JWT Secret Key generated and saved to %s\n", *envFilePath)
}
