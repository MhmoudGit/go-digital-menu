package controllers

import (
	"net/http"
	"strconv"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/go-chi/chi/v5"
)

func AllCategories(w http.ResponseWriter, r *http.Request) {
	data, err := h.GetCategories(database.Db, 1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	h.JsonMarshal(data, w)
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
	h.JsonMarshal(data, w)
}

func PostCategory(w http.ResponseWriter, r *http.Request) {}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {}
