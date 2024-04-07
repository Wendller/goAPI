package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Wendller/goexpert/apis/configs"
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	serverDir := filepath.Join(currentDir, "cmd", "server")

	err = os.Chdir(serverDir)
	if err != nil {
		panic(err)
	}

	_, err = configs.LoadConfig(serverDir)
	if err != nil {
		panic(err)
	}
	defer os.Chdir(currentDir)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.User{}, &entities.Product{})

	repositories := configs.InitializeRepositories(db)
	handlers := configs.InitializeHandlers(repositories)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/products", handlers.ProductHandler.CreateProduct)
	router.Get("/products/{id}", handlers.ProductHandler.GetProduct)
	router.Put("/products/{id}", handlers.ProductHandler.UpdateProduct)

	http.ListenAndServe(":8080", router)
}
