package commands

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type ListProductsCommand struct {
	ProductRepository repositories.ProductRepository
}

func NewListProductsCommand(productRepository repositories.ProductRepository) *ListProductsCommand {
	return &ListProductsCommand{
		ProductRepository: productRepository,
	}
}

func (command *ListProductsCommand) Execute(input *inputs.ListProductsInput) ([]entities.Product, error) {
	products, err := command.ProductRepository.FindMany(input.Page, input.Limit, input.Sort)
	if err != nil {
		return nil, err
	}

	return products, nil
}
