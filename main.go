package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MhmoudGit/go-digital-menu/database"
	routes "github.com/MhmoudGit/go-digital-menu/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// database connection
	database.Connect()
	database.AutoMigrateDb()

	// declaring chi mux as r
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(cors.Handler(corsMiddleware))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// api routes
	routes.HomeRoutes(r)
	routes.PlansRoutes(r)
	routes.AuthRoutes(r)
	routes.RestaurantsRoutes(r)
	routes.CategoriesRoutes(r)
	routes.ProductsRoutes(r)
	withGracefulShuDown(r)
}

// Configure CORS middleware
var corsMiddleware = cors.Options{
	AllowedOrigins:   []string{"*"}, // Replace with your frontend's URL
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"*"},
	AllowCredentials: true,
}

func withGracefulShuDown(r *chi.Mux) {
	// listening on port 8000
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	database.Close()
	log.Println("Server gracefully stopped...")
}
