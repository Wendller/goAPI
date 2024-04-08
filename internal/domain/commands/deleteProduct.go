package commands

import (
	customerrors "github.com/Wendller/goexpert/apis/internal/domain/customErrors"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type DeleteProductCommand struct {
	ProductRepository repositories.ProductRepository
}

func NewDeleteProductCommand(productRepository repositories.ProductRepository) *DeleteProductCommand {
	return &DeleteProductCommand{
		ProductRepository: productRepository,
	}
}

func (command *DeleteProductCommand) Execute(input *inputs.DeleteProductInput) error {
	_, err := command.ProductRepository.FindByID(input.Id)
	if err != nil {
		return customerrors.ErrResourceNotFound
	}

	err = command.ProductRepository.Delete(input.Id)
	if err != nil {
		return err
	}

	return nil
}
