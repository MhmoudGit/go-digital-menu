package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
)

func PlansRoutes(r chi.Router) {
	r.Route("/plans", func(r chi.Router) {
		// get request for Plans route that takes a handler function
		r.Get("/", c.AllPlans)
		r.Get("/{id}", c.SinglePlan)
		r.Post("/", c.PostPlan)
		r.Put("/{id}", c.UpdatePlan)
		r.Delete("/{id}", c.DeletePlan)

		// r.Route("/", func(r chi.Router) {
		// 	r.Use(jwtauth.Verifier(c.TokenAuth))
		// 	r.Use(jwtauth.Authenticator)
		// 	r.Post("/", c.PostPlan)
		// 	r.Put("/{id}", c.UpdatePlan)
		// 	r.Delete("/{id}", c.DeletePlan)
		// 	r.Patch("/{id}/image", c.UpdatePlanImage)
		// })
	})
}
