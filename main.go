package main

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/database"
	routes "github.com/MhmoudGit/go-digital-menu/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// database connection
	database.Connect()
	database.AutoMigrateDb()
	defer database.Close()

	// declaring chi mux as r
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// api routes
	routes.HomeRoutes(r)
	routes.CategoriesRoutes(r)
	routes.ProductsRoutes(r)
	routes.AuthRoutes(r)

	// listening on port 8000
	http.ListenAndServe("127.0.0.1:8000", r)
}
