package inputs

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetProductInput struct {
	Id string `json:"id"`
}

func NewGetProductInput(request *http.Request) (*GetProductInput, error) {
	var input GetProductInput

	id := chi.URLParam(request, "id")
	if id == "" {
		return nil, errors.New("id param is empty")
	}

	input.Id = id

	return &input, nil
}
