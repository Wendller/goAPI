package handlers

import gorm_repositories "github.com/Wendller/goexpert/apis/internal/infra/database/repositories/gorm"

type Handlers struct {
	ProductHandler *ProductHandler
	UserHandler    *UserHandler
}

func NewHandlers(repositories *gorm_repositories.Repositories) *Handlers {
	return &Handlers{
		ProductHandler: NewProductHandler(repositories.ProductRepository),
		UserHandler:    NewUserHandler(repositories.UserRepository),
	}
}
