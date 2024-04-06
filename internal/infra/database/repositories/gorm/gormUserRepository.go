package gorm_repositories

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{DB: db}
}

func (u *GORMUserRepository) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *GORMUserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
