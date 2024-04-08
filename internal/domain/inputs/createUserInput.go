package inputs

import (
	"encoding/json"
	"net/http"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserInput(request *http.Request) (*CreateUserInput, error) {
	var input CreateUserInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
