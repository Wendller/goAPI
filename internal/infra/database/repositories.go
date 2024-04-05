package database

import "github.com/Wendller/goexpert/apis/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductRepository interface {
	Create(product *entity.Product) error
	FindMany(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
