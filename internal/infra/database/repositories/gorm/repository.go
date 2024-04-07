package gorm_repositories

import "gorm.io/gorm"

type Repositories struct {
	ProductRepository *GORMProductRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductRepository: NewProductRepository(db),
	}
}
