package inputs

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UpdateProductInput struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewUpdateProductInput(request *http.Request) (*UpdateProductInput, error) {
	var input UpdateProductInput

	id := chi.URLParam(request, "id")
	if id == "" {
		return nil, errors.New("id param is empty")
	}

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	input.Id = id

	return &input, nil
}
