package user

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// PasswordResetToken stores details for resetting passwords
type PasswordResetToken struct {
	Token     string
	UserEmail string
	ExpiresAt time.Time
}

func GenerateJWTToken(username string) (string, error) {
	expirationTIme := time.Now().Add(24 * time.Hour).Unix()

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTIme,
	}

	// Create the token using the HS256 signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := os.Getenv("JWT_SECRET")

	if jwtSecretKey == "" {
		log.Println("JWT_SECRET environment variable is not set")
		return "", errors.New("JWT_SECRET is not set")
	}

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain text password
func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateResetToken generates a secure random token
func GenerateResetToken() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Stores the reset token in the in-memory map.
func StoreResetToken(token, email string) {
	passwordResetTokens[token] = &PasswordResetToken{
		Token:     token,
		UserEmail: email,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
}

// Validates the reset token, checking its existence and expiration.
func ValidateResetToken(token string) (*PasswordResetToken, error) {
	resetToken, exists := passwordResetTokens[token]
	if !exists || time.Now().After(resetToken.ExpiresAt) {
		return nil, errors.New("invalid or expired token")
	}
	return resetToken, nil
}

// Removes the token from the in-memory map after it has been used.
func InvalidateResetToken(token string) {
	delete(passwordResetTokens, token)
}
