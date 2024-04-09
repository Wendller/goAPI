package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Wendller/goexpert/goAPI/internal/domain/commands"
	customerrors "github.com/Wendller/goexpert/goAPI/internal/domain/customErrors"
	"github.com/Wendller/goexpert/goAPI/internal/domain/inputs"
	"github.com/Wendller/goexpert/goAPI/internal/domain/repositories"
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

func (handler *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewListProductsInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	listProductsCommand := commands.NewListProductsCommand(handler.ProductRepository)

	products, err := listProductsCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
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
		if err == customerrors.ErrResourceNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewDeleteProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deleteProductCommand := commands.NewDeleteProductCommand(handler.ProductRepository)

	err = deleteProductCommand.Execute(input)
	if err != nil {
		if err == customerrors.ErrResourceNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
