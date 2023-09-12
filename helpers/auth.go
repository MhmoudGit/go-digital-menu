package helpers

import (
	"math"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
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

func GenerateAccessToken(providerID uint, tokenAuth *jwtauth.JWTAuth) (string, error) {
	// TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"id":  providerID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	return tokenString, nil
}

func GetProviderIdClaim(r *http.Request) uint {
	_, claims, _ := jwtauth.FromContext(r.Context())
	var iduint uint
	id, ok := claims["id"].(float64)
	if ok {
		iduint = uint(math.Floor(id))
		return iduint
	}
	return 0
}
