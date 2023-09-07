package routers

import (
	"net/http"

	"github.com/MhmoudGit/go-digital-menu/controllers"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CategoriesRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/categories", func(r chi.Router) {
		// get request for home route that takes a handler function
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			controllers.AllCategories(w, r, db)
		})
	})
}
