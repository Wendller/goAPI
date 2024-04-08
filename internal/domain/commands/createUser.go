package commands

import (
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type CreateUserCommand struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserCommand(userRepository repositories.UserRepository) *CreateUserCommand {
	return &CreateUserCommand{
		UserRepository: userRepository,
	}
}

func (command *CreateUserCommand) Execute(input *inputs.CreateUserInput) error {
	user, err := entities.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return err
	}

	err = command.UserRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
