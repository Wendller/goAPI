package routes

import (
	"github.com/Wendller/goexpert/apis/internal/infra/web/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux, handlers *handlers.Handlers) {
	router.Post("/products", handlers.ProductHandler.CreateProduct)
	router.Get("/products", handlers.ProductHandler.ListProducts)
	router.Get("/products/{id}", handlers.ProductHandler.GetProduct)
	router.Put("/products/{id}", handlers.ProductHandler.UpdateProduct)
	router.Delete("/products/{id}", handlers.ProductHandler.DeleteProduct)

	router.Post("/users", handlers.UserHandler.CreateUser)
}
