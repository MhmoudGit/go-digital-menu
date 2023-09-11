package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// verify Provider password
func AuthinticateProvider(db *gorm.DB, email, password string) (bool, error) {
	Provider, err := getProviderByEmail(db, email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(Provider.Password), []byte(password))
	if err != nil {
		// Handle error, e.g. return authentication failure
		return false, err
	}
	// Passwords match
	return true, nil
}

func GenerateAccessToken(providerID int, role string) (string, error) {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	// Define the claims for the token
	claims := jwt.MapClaims{}
	claims["id"] = providerID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate the token using HMAC SHA256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
