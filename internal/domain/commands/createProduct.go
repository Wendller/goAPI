package commands

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type CreateProductCommand struct {
	ProductRepository repositories.ProductRepository
}

func NewCreateProductCommand(productRepository repositories.ProductRepository) *CreateProductCommand {
	return &CreateProductCommand{
		ProductRepository: productRepository,
	}
}

func (command *CreateProductCommand) Execute(input *inputs.CreateProductInput) error {
	product, err := entities.NewProduct(input.Name, input.Price)
	if err != nil {
		return err
	}

	err = command.ProductRepository.Create(product)
	if err != nil {
		return err
	}

	return nil
}
