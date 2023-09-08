package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonMarshal(data any, w http.ResponseWriter) {
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

func JsonDecoder(from io.Reader, to interface{}, w http.ResponseWriter) error {
	// Decode the JSON request body into the struct
	decoder := json.NewDecoder(from)
	if err := decoder.Decode(&to); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		http.Error(w, "Failed to parse JSON request body", http.StatusUnprocessableEntity)
		return err
	}
	return nil
}
