package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// verify Provider password
func AuthinticateProvider(db *gorm.DB, email, password string) (bool, error) {
	Provider, err := GetProviderByEmail(db, email)
	if err != nil {
		return false, err
	}
	err = Provider.VerifyPassword(password)
	if err != nil {
		return false, err
	}
	// Passwords match
	return true, nil
}

func GenerateAccessToken(providerID uint) (string, error) {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	// Define the claims for the token
	claims := jwt.MapClaims{}
	claims["id"] = providerID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate the token using HMAC SHA256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// decrypt the token
func Decrypt(token *jwt.Token) (interface{}, error) {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(jwtSecret), nil
}

// check token and remove Bearer
func CheckToken(token string, w http.ResponseWriter) string {
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Credentials"))
		return ""
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	return token
}
