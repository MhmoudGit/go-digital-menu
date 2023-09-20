package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"github.com/MhmoudGit/go-digital-menu/models"
	u "github.com/MhmoudGit/go-digital-menu/utils"
	"github.com/go-chi/chi/v5"
)

func AllPlans(w http.ResponseWriter, r *http.Request) {
	data, err := h.GetPlans(database.Db)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	// json response
	u.JsonMarshal(data, w)
}

func SinglePlan(w http.ResponseWriter, r *http.Request) {
	// Convert the string to a uint
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	data, err := h.GetPlan(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	u.JsonMarshal(data, w)
}

func PostPlan(w http.ResponseWriter, r *http.Request) {
	var validPlan models.Plan
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validPlan, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	// store the struct data into the database
	err = h.CreatePlan(database.Db, &validPlan)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	u.JsonMarshal(&validPlan, w)
}

func UpdatePlan(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	var validPlan models.Plan
	// store the json request body into my struct
	err := u.JsonDecoder(r.Body, &validPlan, w)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	validPlan.ID = id
	err = h.UpdatePlan(database.Db, &validPlan, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.JsonMarshal(&validPlan, w)
}

func DeletePlan(w http.ResponseWriter, r *http.Request) {
	id := u.ParseUint64(w, chi.URLParam(r, "id"))
	err := h.DeletePlan(database.Db, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusAccepted)
}
