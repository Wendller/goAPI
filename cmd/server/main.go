package main

import (
	"net/http"

	"github.com/Wendller/goexpert/apis/configs"
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/infra/web/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	router.Use(middleware.Logger)

	routes.SetupRoutes(router, handlers)

	http.ListenAndServe(":8080", router)
}
