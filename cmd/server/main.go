package main

import (
	"net/http"

	"github.com/Wendller/goexpert/goAPI/configs"
	"github.com/Wendller/goexpert/goAPI/internal/domain/entities"
	"github.com/Wendller/goexpert/goAPI/internal/infra/web/routes"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title GO API
// @version 1.0
// @description Product API with JWT authentication
// @contact.name Wendler

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.User{}, &entities.Product{})

	repositories := configs.InitializeRepositories(db)
	handlers := configs.InitializeHandlers(repositories)

	router := chi.NewRouter()
	routes.SetupRoutes(router, handlers)

	http.ListenAndServe(":8080", router)
}
