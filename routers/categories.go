package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
)

func CategoriesRoutes(r chi.Router) {
	r.Route("/categories", func(r chi.Router) {

		// get request for categories route that takes a handler function
		r.Get("/", c.AllCategories)
		r.Post("/", c.PostCategory)
		r.Get("/{id}", c.SingleCategory)
		r.Put("/{id}", c.UpdateCategory)
		r.Delete("/{id}", c.DeleteCategory)
	})
}
