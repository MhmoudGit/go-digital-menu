package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	h "github.com/MhmoudGit/go-digital-menu/helpers"
)

func AllCategories(w http.ResponseWriter, r *http.Request) {
	data, err := h.GetCategories(database.Db, 1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	/// json marshaling data
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type and send the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
