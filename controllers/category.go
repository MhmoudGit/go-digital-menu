package controllers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
)

func AllCategories(w http.ResponseWriter, r *http.Request) {
	data, err := h.GetCategories(database.Db, 1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	h.JsonMarshal(data, w)
}
