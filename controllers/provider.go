package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func SingleProvider(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	// Convert the string to a uint
	data, err := h.GetProvider(database.Db, id)
	if err != nil {
		u.NotFound(w)
		return
	}
	u.JsonMarshal(data, w)
}

func UpdateProvider(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	var validProvider models.UpdateProvider
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validProvider, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = h.UpdateProvider(database.Db, &validProvider, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.JsonMarshal(&validProvider, w)
}

func UpdateProviderImage(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	u.ParseMultipartForm(w, r)
	validProviderImage := models.UpdateProviderImage{
		Image: u.UploadFile(w, r, "image"),
	}
	if validProviderImage.Image != "" {
		// store the struct data into the database
		err := h.UpdateProviderImage(database.Db, &validProviderImage, id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validProviderImage, w)
}

func DeleteProvider(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	err := h.DeleteProvider(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusAccepted)
}
