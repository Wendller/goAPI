package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Wendller/goexpert/apis/internal/domain/commands"
	customerrors "github.com/Wendller/goexpert/apis/internal/domain/customErrors"
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

func (handler *UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewSignInUserInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	signInUserCommand := commands.NewSignInUserCommand(handler.UserRepository)

	accessToken, err := signInUserCommand.Execute(input)
	if err != nil {
		if err == customerrors.ErrWrongCredentials {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
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
