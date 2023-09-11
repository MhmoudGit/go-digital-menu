package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {

		r.Post("/register", c.RegisterHandler)
		r.Post("/login", c.LoginHandler)
	})
}
