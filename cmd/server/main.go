package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Wendller/goexpert/apis/configs"
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	gorm_repositories "github.com/Wendller/goexpert/apis/internal/infra/database/repositories/gorm"
	"github.com/Wendller/goexpert/apis/internal/infra/web/handlers"
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

	productRepository := gorm_repositories.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
