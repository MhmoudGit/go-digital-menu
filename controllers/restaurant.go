package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	// "github.com/MhmoudGit/go-digital-menu/models"
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

// func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
// 	id := u.ParseUint64(w, chi.URLParam(r, "id"))
// 	var validRestaurant models.UpdateRestaurant
// 	// store the json request body into my struct
// 	err := u.JsonDecoder(r.Body, &validRestaurant, w)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		return
// 	}
// 	err = h.UpdateRestaurant(database.Db, &validRestaurant, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		return
// 	}
// 	u.JsonMarshal(&validRestaurant, w)
// }

// func UpdateRestaurantImage(w http.ResponseWriter, r *http.Request) {
// 	id := u.ParseUint64(w, chi.URLParam(r, "id"))
// 	u.ParseMultipartForm(w, r)
// 	validRestaurantImage := models.UpdateRestaurantImage{
// 		Image: u.UploadFile(w, r, "image"),
// 	}
// 	if validRestaurantImage.Image != "" {
// 		// store the struct data into the database
// 		err := h.UpdateRestaurantImage(database.Db, &validRestaurantImage, id)
// 		if err != nil {
// 			w.WriteHeader(http.StatusUnprocessableEntity)
// 		}
// 	}
// 	u.JsonMarshal(&validRestaurantImage, w)
// }

// func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
// 	id := u.ParseUint64(w, chi.URLParam(r, "id"))
// 	err := h.DeleteRestaurant(database.Db, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 	}
// 	w.WriteHeader(http.StatusAccepted)
// }
