package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomeRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		// get request for home route that takes a handler function
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"hello": "world"}`))
		})

		// Define a route to serve the HTML file
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./views/openapi.html") // Replace "index.html" with your HTML file path
		})

		// Define a route to serve the HTML file
		r.Get("/openapi", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./openapi.yaml") // Replace "index.html" with your HTML file path
		})
	})
}
