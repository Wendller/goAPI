package commands

import (
	customerrors "github.com/Wendller/goexpert/apis/internal/domain/customErrors"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
	chyptograhpy_adapter "github.com/Wendller/goexpert/apis/internal/infra/cryptography"
)

type SignInUserCommand struct {
	UserRepository repositories.UserRepository
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

func NewSignInUserCommand(userRepository repositories.UserRepository) *SignInUserCommand {
	return &SignInUserCommand{
		UserRepository: userRepository,
	}
}

func (command *SignInUserCommand) Execute(input *inputs.SignInUserInput) (*SignInResponse, error) {
	user, err := command.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, customerrors.ErrWrongCredentials
	}

	bcryptHasher := chyptograhpy_adapter.NewBcryptHasher()
	isValidPassword := bcryptHasher.Compare(user.Password, input.Password)

	if !isValidPassword {
		return nil, customerrors.ErrWrongCredentials
	}

	jwtAuthEncrypter := chyptograhpy_adapter.NewJWTEncrypter()
	accessToken := jwtAuthEncrypter.Encrypt(map[string]interface{}{"userID": user.ID})

	return &SignInResponse{
		AccessToken: accessToken,
	}, nil
}
