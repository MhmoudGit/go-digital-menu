package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func SingleRestaurant(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	// Convert the string to a uint
	data, err := h.GetRestaurant(database.Db, id)
	if err != nil {
		u.NotFound(w)
		return
	}
	u.JsonMarshal(data, w)
}

func PostRestaurant(w http.ResponseWriter, r *http.Request) {
	userId := h.GetUserIdClaim(r)
	u.ParseMultipartForm(w, r)
	newRestaurant := models.NewRestaurant(
		userId,
		r.FormValue("name"),
		r.FormValue("enName"),
		u.UploadFile(w, r, "image"),
		u.UploadFile(w, r, "theme"),
		u.UploadFile(w, r, "cover"),
		r.FormValue("whatsapp"),
		r.FormValue("url"),
		r.FormValue("googleMap"),
		u.ParseTime(r.FormValue("openedFrom")),
		u.ParseTime(r.FormValue("openedTo")),
		u.Parseint(w, r.FormValue("tables")),
	)
	// store the struct data into the database
	err := h.CreateRestaurant(database.Db, newRestaurant)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return 
	}
	u.JsonMarshal(&newRestaurant.ID, w)
}

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	userId := h.GetUserIdClaim(r)
	var validRestaurant models.UpdateRestaurant
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validRestaurant, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = h.UpdateRestaurant(database.Db, &validRestaurant, id, userId)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.JsonMarshal(&validRestaurant, w)
}

func UpdateRestaurantImage(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	u.ParseMultipartForm(w, r)
	validRestaurantImage := models.UpdateRestaurantImage{
		Image: u.UploadFile(w, r, "image"),
	}
	if validRestaurantImage.Image != "" {
		// store the struct data into the database
		err := h.UpdateRestaurantImage(database.Db, &validRestaurantImage, id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validRestaurantImage, w)
}

func UpdateRestaurantCover(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	u.ParseMultipartForm(w, r)
	validRestaurantCover := models.UpdateRestaurantCover{
		Cover: u.UploadFile(w, r, "cover"),
	}
	if validRestaurantCover.Cover != "" {
		// store the struct data into the database
		err := h.UpdateRestaurantCover(database.Db, &validRestaurantCover, id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validRestaurantCover, w)
}

func UpdateRestaurantTheme(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	u.ParseMultipartForm(w, r)
	validRestaurantTheme := models.UpdateRestaurantTheme{
		Theme: u.UploadFile(w, r, "theme"),
	}
	if validRestaurantTheme.Theme != "" {
		// store the struct data into the database
		err := h.UpdateRestaurantTheme(database.Db, &validRestaurantTheme, id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validRestaurantTheme, w)
}

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	err := h.DeleteRestaurant(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusAccepted)
}
