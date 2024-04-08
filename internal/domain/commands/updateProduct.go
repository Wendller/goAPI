package commands

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
	"github.com/Wendller/goexpert/apis/pkg/entity"
)

type UpdateProductCommand struct {
	ProductRepository repositories.ProductRepository
}

func NewUpdateProductCommand(productRepository repositories.ProductRepository) *UpdateProductCommand {
	return &UpdateProductCommand{
		ProductRepository: productRepository,
	}
}

func (command *UpdateProductCommand) Execute(input *inputs.UpdateProductInput) error {
	_, err := command.ProductRepository.FindByID(input.Id)
	if err != nil {
		return err
	}

	id, _ := entity.ParseID(input.Id)

	product := entities.Product{
		ID:    id,
		Name:  input.Name,
		Price: input.Price,
	}

	err = command.ProductRepository.Update(&product)
	if err != nil {
		return err
	}

	return nil
}
