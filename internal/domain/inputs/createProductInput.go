package inputs

import (
	"encoding/json"
	"net/http"
)

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewCreateProductInput(request *http.Request) (*CreateProductInput, error) {
	var input CreateProductInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
