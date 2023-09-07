package controllers

import (
	"encoding/json"
	"net/http"

	h "github.com/MhmoudGit/go-digital-menu/helpers"
	"gorm.io/gorm"
)

func AllCategories(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	data, err := h.GetCategories(db, 1)
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
