package handlers

import (
	"encoding/json"
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
	input, err := inputs.NewCreateProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createProductCommand := commands.NewCreateProductCommand(handler.ProductRepository)

	err = createProductCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewGetProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getProductCommand := commands.NewGetProductCommand(handler.ProductRepository)

	product, err := getProductCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewUpdateProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateProductCommand := commands.NewUpdateProductCommand(handler.ProductRepository)

	err = updateProductCommand.Execute(input)
	if err != nil {
		if err.Error() == "raw not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
