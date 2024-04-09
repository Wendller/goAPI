package repositories

import "github.com/Wendller/goexpert/goAPI/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
