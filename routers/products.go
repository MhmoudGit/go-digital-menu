package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
)

func ProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {

		// get request for categories route that takes a handler function
		r.Get("/", c.AllProducts)
		r.Post("/", c.PostProduct)
		r.Get("/{id}", c.SingleProduct)
		r.Put("/{id}", c.UpdateProduct)
		r.Delete("/{id}", c.DeleteProduct)
		r.Patch("/{id}/image", c.UpdateProductImage)
	})
}
