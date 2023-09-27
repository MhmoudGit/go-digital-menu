package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func AuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", c.RegisterHandler)
		r.Post("/login", c.LoginHandler)
		r.Get("/refresh", c.RefreshTokenHandler)
		r.Get("/verify-email/{id}", c.VerifyEmail)

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Patch("/change-password", c.ChangePassword)
		})
	})
}
