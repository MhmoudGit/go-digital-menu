package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RestaurantsRoutes(r chi.Router) {
	r.Route("/restaurants", func(r chi.Router) {

		// get request for Restaurants route that takes a handler function
		r.Get("/{id}", c.SingleRestaurant)

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", c.PostRestaurant)
			r.Put("/{id}", c.UpdateRestaurant)
			r.Delete("/{id}", c.DeleteRestaurant)
			r.Patch("/{id}/image", c.UpdateRestaurantImage)
			r.Patch("/{id}/cover", c.UpdateRestaurantCover)
			r.Patch("/{id}/theme", c.UpdateRestaurantTheme)
		})
	})
}
