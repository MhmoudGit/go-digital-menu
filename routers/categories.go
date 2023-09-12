package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func CategoriesRoutes(r chi.Router) {
	r.Route("/categories", func(r chi.Router) {
		// get request for categories route that takes a handler function
		r.Get("/", c.AllCategories)
		r.Get("/{id}", c.SingleCategory)

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", c.PostCategory)
			r.Put("/{id}", c.UpdateCategory)
			r.Delete("/{id}", c.DeleteCategory)
			r.Patch("/{id}/image", c.UpdateCategoryImage)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", c.AllCategoriesPrivate)
		})
	})
}
