package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomeRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		// get request for home route that takes a handler function
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./views/main.html")
		})

		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./views/openapi.html")
		})

		// uploads folder
		uploadsDir := "./uploads" // Change this to the path of your uploads folder
		r.Get("/uploads/*", func(w http.ResponseWriter, r *http.Request) {
			http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadsDir))).ServeHTTP(w, r)
		})
	})
}
