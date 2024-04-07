package commands

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type GetProductCommand struct {
	ProductRepository repositories.ProductRepository
}

func NewGetProductCommand(productRepository repositories.ProductRepository) *GetProductCommand {
	return &GetProductCommand{
		ProductRepository: productRepository,
	}
}

func (command *GetProductCommand) Execute(input *inputs.GetProductInput) (*entities.Product, error) {
	product, err := command.ProductRepository.FindByID(input.Id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
