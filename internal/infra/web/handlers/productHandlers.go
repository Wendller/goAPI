package handlers

import (
	"net/http"

	"github.com/Wendller/goexpert/apis/internal/domain/commands"
	"github.com/Wendller/goexpert/apis/internal/domain/inputs"
	"github.com/Wendller/goexpert/apis/internal/domain/repositories"
)

type ProductHandler struct {
	ProductRepository repositories.ProductRepository
}

func NewProductHandler(productRepository repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepository: productRepository,
	}
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewCreateProductInput(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	command := commands.NewCreateProductCommand(handler.ProductRepository)

	err = command.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
