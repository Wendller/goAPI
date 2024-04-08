package inputs

import (
	"net/http"

	customerrors "github.com/Wendller/goexpert/apis/internal/domain/customErrors"
	"github.com/go-chi/chi/v5"
)

type DeleteProductInput struct {
	Id string `json:"id"`
}

func NewDeleteProductInput(request *http.Request) (*DeleteProductInput, error) {
	var input DeleteProductInput

	id := chi.URLParam(request, "id")
	if id == "" {
		return nil, customerrors.ErrEmptyID
	}

	input.Id = id

	return &input, nil
}
