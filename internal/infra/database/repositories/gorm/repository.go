package gorm_repositories

import "gorm.io/gorm"

type Repositories struct {
	ProductRepository *GORMProductRepository
	UserRepository    *GORMUserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductRepository: NewProductRepository(db),
		UserRepository:    NewUserRepository(db),
	}
}
