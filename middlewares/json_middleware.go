package middlewares

import (
	"encoding/json"
	"net/http"
)

func JSONMiddleware(schema interface{}, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(schema)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		// Set the content type and send the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		next.ServeHTTP(w, r) // Call the next middleware or handler
	})
}
