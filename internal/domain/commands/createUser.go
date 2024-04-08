package commands

import (
	customerrors "github.com/Wendller/goexpert/apis/internal/domain/customErrors"
	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
	chyptograhpy_hasher "github.com/Wendller/goexpert/apis/internal/infra/cryptography"
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
	userWithSameCredential, _ := command.UserRepository.FindByEmail(input.Email)
	if userWithSameCredential != nil {
		return customerrors.ErrUserAlreadyExists
	}

	bcryptHasher := chyptograhpy_hasher.NewBcryptHasher()

	passwordHash, err := bcryptHasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user, err := entities.NewUser(input.Name, input.Email, string(passwordHash))
	if err != nil {
		return err
	}

	err = command.UserRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
