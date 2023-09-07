package routers

import (
	"github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
)

func CategoriesRoutes(r chi.Router) {
	r.Route("/categories", func(r chi.Router) {

		// get request for categories route that takes a handler function
		r.Get("/", controllers.AllCategories)
	})
}
