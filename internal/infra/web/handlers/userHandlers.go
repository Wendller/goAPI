package handlers

import (
	"net/http"

	"github.com/Wendller/goexpert/apis/internal/domain/commands"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewCreateUserInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createUserCommand := commands.NewCreateUserCommand(handler.UserRepository)

	err = createUserCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
