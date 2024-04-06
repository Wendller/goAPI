package inputs

import (
	"encoding/json"
	"io"
)

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewCreateProductInput(params io.ReadCloser) (*CreateProductInput, error) {
	var input CreateProductInput

	err := json.NewDecoder(params).Decode(&input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
