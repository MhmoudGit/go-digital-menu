package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func AllCategories(w http.ResponseWriter, r *http.Request) {
	resQueryParam := u.ParseUint64(w, r.URL.Query().Get("resid"))
	data, err := h.GetCategories(database.Db, resQueryParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	// json response
	u.JsonMarshal(data, w)
}

func SingleCategory(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	// Convert the string to a uint
	data, err := h.GetCategory(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	u.JsonMarshal(data, w)
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	resId := h.GetResIdClaim(r)
	// Parse the form data, including the uploaded file
	u.ParseMultipartForm(w, r)
	validCategory := models.PostCategory{
		Name:         r.FormValue("name"),
		EnName:       r.FormValue("enName"),
		Logo:         u.UploadFile(w, r, "logo"),
		RestaurantID: resId,
	}
	// store the struct data into the database
	err := h.CreateCategory(database.Db, &validCategory)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	u.JsonMarshal(&validCategory, w)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	resId := h.GetResIdClaim(r)
	var validCategory models.UpdateCategory
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validCategory, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = h.UpdateCategory(database.Db, &validCategory, id, resId)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	u.JsonMarshal(&validCategory, w)
}

func UpdateCategoryImage(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	resId := h.GetResIdClaim(r)
	u.ParseMultipartForm(w, r)
	validCategoryImage := models.UpdateCategoryImage{
		Logo: u.UploadFile(w, r, "logo"),
	}
	if validCategoryImage.Logo != "" {
		// store the struct data into the database
		err := h.UpdateCategoryImage(database.Db, &validCategoryImage, id, resId)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validCategoryImage, w)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	resId := h.GetResIdClaim(r)
	err := h.DeleteCategory(database.Db, id, resId)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusAccepted)
}
