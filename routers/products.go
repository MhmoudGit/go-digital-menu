package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {

		// get request for products route that takes a handler function
		r.Get("/", c.AllProducts)
		r.Get("/{id}", c.SingleProduct)

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", c.PostProduct)
			r.Put("/{id}", c.UpdateProduct)
			r.Delete("/{id}", c.DeleteProduct)
			r.Patch("/{id}/image", c.UpdateProductImage)
		})
	})
}
