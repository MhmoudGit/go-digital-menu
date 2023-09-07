package main

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.Connect()
	database.AutoMigrateDb()

	// declaring chi mux as r
	r := chi.NewRouter()
	// using chi logger middleware
	r.Use(middleware.Logger)

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

	// listening on port 8000
	http.ListenAndServe("127.0.0.1:8000", r)
}
