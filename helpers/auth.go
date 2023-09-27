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

func GenerateToken(userId,resID uint, tokenAuth *jwtauth.JWTAuth, expireTime time.Duration) (string, error) {
	// TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"userId": userId,
		"id":  resID,
		"exp": time.Now().Add(time.Hour * expireTime).Unix(),
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

func GetUserIdClaim(r *http.Request) uint {
	_, claims, _ := jwtauth.FromContext(r.Context())
	var userIduint uint
	id, ok := claims["userId"].(float64)
	if ok {
		userIduint = uint(math.Floor(id))
		return userIduint
	}
	return 0
}

func SetCookies(userId,resID uint, tokenAuth *jwtauth.JWTAuth, w http.ResponseWriter) {
	refreshToken, err := GenerateToken(userId,resID, tokenAuth, 24)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}
	// Create a new cookie
	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   86400,
		HttpOnly: true,
	}

	// Set the cookie in the response
	http.SetCookie(w, cookie)
}
