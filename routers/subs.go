package routers

import (
	c "github.com/MhmoudGit/go-digital-menu/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func SubsRoutes(r chi.Router) {
	r.Route("/subscriptions", func(r chi.Router) {

		r.Route("/", func(r chi.Router) {
			r.Use(jwtauth.Verifier(c.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Patch("/verify-payment", c.VerifyPayment)
		})
	})
}
