package helpers

import (
	"encoding/json"
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
