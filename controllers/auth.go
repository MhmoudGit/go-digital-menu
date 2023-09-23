package controllers

import (
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte(getSecret()), nil)

func getSecret() string {
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
	if user.IsActive && userAuth {
		// Generate an access token for the authenticated User
		accessToken, err := h.GenerateAccessToken(user.Restaurant.ID, TokenAuth)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}
