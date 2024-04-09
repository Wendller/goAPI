package routes

import (
	_ "github.com/Wendller/goexpert/goAPI/docs"
	"github.com/Wendller/goexpert/goAPI/internal/infra/auth"
	"github.com/Wendller/goexpert/goAPI/internal/infra/web/handlers"
	"github.com/Wendller/goexpert/goAPI/internal/infra/web/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(router *chi.Mux, handlers *handlers.Handlers) {
	JWTAuthConfig := auth.NewJWTAuthConfig()
	tokenAuth := JWTAuthConfig.JWT

	router.Use(middlewares.LogRequest)
	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", handlers.ProductHandler.CreateProduct)
		r.Get("/", handlers.ProductHandler.ListProducts)
		r.Get("/{id}", handlers.ProductHandler.GetProduct)
		r.Put("/{id}", handlers.ProductHandler.UpdateProduct)
		r.Delete("/{id}", handlers.ProductHandler.DeleteProduct)
	})

	router.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.UserHandler.CreateUser)
		r.Post("/sessions", handlers.UserHandler.SignInUser)
	})

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
}
