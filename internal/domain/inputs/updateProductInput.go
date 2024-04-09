package inputs

import (
	"encoding/json"
	"net/http"

	customerrors "github.com/Wendller/goexpert/goAPI/internal/domain/customErrors"
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
		return nil, customerrors.ErrEmptyID
	}

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	input.Id = id

	return &input, nil
}
