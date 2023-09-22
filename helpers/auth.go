package helpers

import (
	"math"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

// verify User password
func AuthinticateUser(db *gorm.DB, email, password string) (bool, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return false, err
	}
	err = user.VerifyPassword(password)
	if err != nil {
		return false, err
	}
	// Passwords match
	return true, nil
}

func GenerateAccessToken(resID uint, tokenAuth *jwtauth.JWTAuth) (string, error) {
	// TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"id":  resID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	return tokenString, nil
}

func GetResIdClaim(r *http.Request) uint {
	_, claims, _ := jwtauth.FromContext(r.Context())
	var iduint uint
	id, ok := claims["id"].(float64)
	if ok {
		iduint = uint(math.Floor(id))
		return iduint
	}
	return 0
}
