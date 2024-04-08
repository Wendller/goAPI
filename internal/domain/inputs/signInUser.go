package inputs

import (
	"encoding/json"
	"net/http"
)

type SignInUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignInUserInput(request *http.Request) (*SignInUserInput, error) {
	var input SignInUserInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
