package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte(getSecret()), nil)

func getSecret() string {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtSecret
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginForm models.Login
	// Authenticate the User and retrieve the User ID
	// Parse the request body
	err := u.JsonDecoder(r.Body, &loginForm, w)
	if err != nil {
		http.Error(w, "wrong email or password", http.StatusBadRequest)
		return
	}
	// authinticate User
	userAuth, err := h.AuthinticateUser(database.Db, loginForm.Email, loginForm.Password)
	if err != nil {
		http.Error(w, "wrong email or password", http.StatusUnauthorized)
		return
	}
	user, err := h.GetUserByEmail(database.Db, loginForm.Email)
	if err != nil {
		http.Error(w, "wrong email or password", http.StatusUnauthorized)
		return
	}
	if user.IsActive && userAuth && user.IsVerified {
		// Generate an access token for the authenticated User
		accessToken, err := h.GenerateToken(user.ID, user.Restaurant.ID, TokenAuth, 1)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}
		h.SetCookies(user.ID, user.Restaurant.ID, TokenAuth, w)
		// Return the access token to the client
		w.Header().Set("Authorization", "Bearer "+accessToken)
		u.JsonMarshal(&user, w)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("wrong email or password"))
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	// Parse the form data, including the uploaded file
	u.ParseMultipartForm(w, r)
	newUser := models.NewUser(
		u.Parseint(w, r.FormValue("duration")),
		u.ParseUint64(w, r.FormValue("planId")),
		r.FormValue("email"),
		r.FormValue("password"),
		r.FormValue("restaurantName"),
		r.FormValue("phone"),
	)
	// check if email exists in db
	_, err := h.GetUserByEmail(database.Db, newUser.Email)
	if err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "Try again", http.StatusBadRequest)
		w.Write([]byte("Try again"))
		return
	}
	err = h.CreateUser(database.Db, newUser)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("going to send email")
	wg.Add(1)
	go func() {
		body := fmt.Sprintf("Body: click on the link to verify your email: http://localhost:8000/auth/verify-email/%v", newUser.ID)
		defer wg.Done()
		u.SendEmail(body, newUser.Email)
	}()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
	wg.Wait()
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	oldRefreshToken := jwtauth.TokenFromCookie(r)
	fmt.Println(oldRefreshToken)
	validToken, err := jwtauth.VerifyToken(TokenAuth, oldRefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	claims := validToken.PrivateClaims()
	resId := int(claims["id"].(float64))
	userId := int(claims["userId"].(float64))
	user, err := h.GetUser(database.Db, uint(userId))
	if err != nil {
		http.Error(w, "not found", http.StatusUnauthorized)
		return
	}
	if user.Restaurant.ID != uint(resId) {
		http.Error(w, "not found", http.StatusUnauthorized)
		return
	}
	accessToken, err := h.GenerateToken(uint(userId), uint(resId), TokenAuth, time.Hour*1)
	if err != nil {
		http.Error(w, "failed to generate access token", http.StatusInternalServerError)
		return
	}
	// Return the refresh token to the client
	u.JsonMarshal(&accessToken, w)
	w.WriteHeader(http.StatusOK)
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	user, err := h.GetUser(database.Db, id)
	if err != nil {
		http.Error(w, "failed to verify email", http.StatusBadRequest)
		return
	}
	if user.IsVerified {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User email is already verified"))
		return
	}
	user.IsVerified = true
	h.UpdateUser(database.Db, &user, user.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User email verified successfully"))
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	userId := h.GetUserIdClaim(r)
	var changePassword models.ChangePassword
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &changePassword, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	user, err := h.GetUser(database.Db, userId)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	err = user.VerifyPassword(changePassword.OldPassword)
	if err != nil {
		http.Error(w, "try again", http.StatusBadRequest)
		return
	}
	err = user.HashPassword(changePassword.NewPassword)
	if err != nil {
		http.Error(w, "error hashing new password", http.StatusBadRequest)
		return
	}
	h.UpdateUser(database.Db, &user, user.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("password changed successfully"))
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var user models.User
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &user.Email, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	user, err = h.GetUserByEmail(database.Db, user.Email)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	tempToken, err := h.GenerateToken(uint(user.ID), uint(user.Restaurant.ID), TokenAuth, time.Minute*5)
	if err != nil {
		http.Error(w, "failed to generate access token", http.StatusInternalServerError)
		return
	}
	wg.Add(1)
	go func() {
		body := fmt.Sprintf("click on the link to verify your email: http://localhost:8000/auth/reset-password/%v", tempToken)
		defer wg.Done()
		u.SendEmail(body, user.Email)
	}()
	wg.Wait()
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	// get token from uri
	tempToken := chi.URLParam(r, "token")
	var resetPassword models.ResetPassword
	// validate the token
	validToken, err := jwtauth.VerifyToken(TokenAuth, tempToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// get the userid from claims
	claims := validToken.PrivateClaims()
	userId := int(claims["userId"].(float64))
	// store the json request body into my struct
	err = u.JsonDecoder(r.Body, &resetPassword, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	//get user
	user, err := h.GetUser(database.Db, uint(userId))
	if err != nil {
		http.Error(w, "not found", http.StatusUnauthorized)
		return
	}
	// store the new password
	if resetPassword.NewPassword != resetPassword.ConfirmPassword {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to change password"))
		return
	}
	err = user.HashPassword(resetPassword.NewPassword)
	if err != nil {
		http.Error(w, "error hashing new password", http.StatusBadRequest)
		return
	}
	h.UpdateUser(database.Db, &user, user.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("password changed successfully"))
}
