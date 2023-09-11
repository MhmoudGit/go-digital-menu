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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	u.ParseMultipartForm(w, r)
	validProvider := models.PostProvider{
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
		Image:       u.UploadFile(w, r, "image"),
		Name:        r.FormValue("name"),
		EnName:      r.FormValue("enName"),
		ServiceType: r.FormValue(("serviceType")),
		Whatsapp:    r.FormValue("whatsapp"),
		Phone:       r.FormValue("phone"),
		Address:     r.FormValue("address"),
		EnAddress:   r.FormValue("enAdress"),
		Facebook:    r.FormValue("facebook"),
		Theme:       r.FormValue("theme"),
		OpenedFrom:  u.ParseTime(r.FormValue("openedFrom")),
		OpenedTo:    u.ParseTime(r.FormValue("openedTo")),
		Url:         r.FormValue("url"),
	}
	// check if email exists in db
	_, err := h.GetProviderByEmail(database.Db, validProvider.Email)
	if err != nil {
		err := h.CreateProvider(database.Db, &validProvider)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "Try again", http.StatusBadRequest)
		return
	}
}
