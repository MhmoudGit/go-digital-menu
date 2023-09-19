package routers

// import (
// 	c "github.com/MhmoudGit/go-digital-menu/controllers"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/jwtauth/v5"
// )

// func ProvidersRoutes(r chi.Router) {
// 	r.Route("/providers", func(r chi.Router) {

// 		// get request for Providers route that takes a handler function
// 		r.Get("/{id}", c.SingleProvider)

// 		r.Route("/", func(r chi.Router) {
// 			r.Use(jwtauth.Verifier(c.TokenAuth))
// 			r.Use(jwtauth.Authenticator)
// 			r.Put("/{id}", c.UpdateProvider)
// 			r.Delete("/{id}", c.DeleteProvider)
// 			r.Patch("/{id}/image", c.UpdateProviderImage)
// 		})
// 	})
// }
