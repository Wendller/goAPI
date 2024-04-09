package repositories

import "github.com/Wendller/goexpert/goAPI/internal/domain/entities"

type ProductRepository interface {
	Create(product *entities.Product) error
	FindMany(page, limit int, sort string) ([]entities.Product, error)
	FindByID(id string) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id string) error
}
