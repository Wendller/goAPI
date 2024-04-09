package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Wendller/goexpert/goAPI/internal/domain/commands"
	customerrors "github.com/Wendller/goexpert/goAPI/internal/domain/customErrors"
	"github.com/Wendller/goexpert/goAPI/internal/domain/inputs"
	"github.com/Wendller/goexpert/goAPI/internal/domain/repositories"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

// SignIn user godoc
// @Summary Login user
// @Description Authenticate user with jwt
// @Tags users
// @Accept json
// @Produce json
// @Param request body inputs.SignInUsetInput true "user credentials"
// @Success 200 {object} commands.SignInResponse
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /users/sessions [post]
func (handler *UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewSignInUserInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body inputs.CreateUserInput true "user request"
// @Success 201
// @Failure 400
// @Failure 500 {object} Error
// @Router /users [post]
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewCreateUserInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	createUserCommand := commands.NewCreateUserCommand(handler.UserRepository)

	err = createUserCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
