package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginForm models.Login
	// Authenticate the provider and retrieve the provider ID
	// Parse the request body
	err := u.JsonDecoder(r.Body, &loginForm, w)
	if err != nil {
		http.Error(w, "wrong email or password", http.StatusBadRequest)
		return
	}
	// authinticate provider
	providerAuth, err := h.AuthinticateProvider(database.Db, loginForm.Email, loginForm.Password)
	if err != nil {
		http.Error(w, "wrong email or password", http.StatusUnauthorized)
		return
	}

	if providerAuth {
		provider, err := h.GetProviderByEmail(database.Db, loginForm.Email)
		if err != nil {
			http.Error(w, "wrong email or password", http.StatusUnauthorized)
			return
		}
		// Generate an access token for the authenticated provider
		accessToken, err := h.GenerateAccessToken(provider.ID)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}
		// Return the access token to the client
		w.Header().Set("Authorization", "Bearer "+accessToken)
		u.JsonMarshal(&provider, w)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("wrong email or password"))
	}

}
