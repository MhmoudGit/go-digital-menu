package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func AllProducts(w http.ResponseWriter, r *http.Request) {
	providerQueryParam := u.ParseUint64(w, r.URL.Query().Get("providerid"))
	categoryQueryParam := u.ParseUint64(w, r.URL.Query().Get("categoryid"))
	data, err := h.GetProducts(database.Db, providerQueryParam, categoryQueryParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	u.JsonMarshal(data, w)
}

func SingleProduct(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	// Convert the string to a uint
	data, err := h.GetProduct(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	u.JsonMarshal(data, w)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	// Parse the form data, including the uploaded file
	u.ParseMultipartForm(w, r)
	validProduct := models.PostProduct{
		Name:       r.FormValue("name"),
		EnName:     r.FormValue("enName"),
		Details:    r.FormValue("details"),
		EnDetails:  r.FormValue("enDetails"),
		Image:      u.UploadFile(w, r, "image"),
		Price:      u.Parseint(w, r.FormValue("price")),
		IsActive:   r.FormValue("isActive") == "true",
		CategoryID: u.ParseUint64(w, r.FormValue("categoryId")),
		ProviderID: u.ParseUint64(w, r.FormValue("providerId")),
	}
	// store the struct data into the database
	err := h.CreateProduct(database.Db, &validProduct)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	u.JsonMarshal(&validProduct, w)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	var validProduct models.UpdateProduct
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validProduct, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = h.UpdateProduct(database.Db, &validProduct, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.JsonMarshal(&validProduct, w)
}

func UpdateProductImage(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	u.ParseMultipartForm(w, r)
	validProductImage := models.UpdateProductImage{
		Image: u.UploadFile(w, r, "image"),
	}
	if validProductImage.Image != "" {
		// store the struct data into the database
		err := h.UpdateProductImage(database.Db, &validProductImage, id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
	u.JsonMarshal(&validProductImage, w)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	err := h.DeleteProduct(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusAccepted)
}
