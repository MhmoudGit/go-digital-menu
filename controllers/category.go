package controllers

import (
	"net/http"
	"strconv"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func AllCategories(w http.ResponseWriter, r *http.Request) {
	data, err := h.GetCategories(database.Db, 1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	u.JsonMarshal(data, w)
}

func SingleCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Convert the string to a uint
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	data, err := h.GetCategory(database.Db, uint(uintId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	u.JsonMarshal(data, w)
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	var validCategory models.PostCategory
	// store the json request body into my struct
	u.JsonDecoder(r.Body, &validCategory, w)
	// store the struct data into the database
	err := h.CreateCategory(database.Db, &validCategory)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	u.JsonMarshal(&validCategory, w)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {}
