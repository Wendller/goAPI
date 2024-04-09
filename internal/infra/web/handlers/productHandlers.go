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

// Create product godoc
// @Summary Create product
// @Description Register a new product
// @Tags products
// @Accept json
// @Produce json
// @Param request body inputs.CreateProductInput true "product request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewCreateProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	createProductCommand := commands.NewCreateProductCommand(handler.ProductRepository)

	err = createProductCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get product godoc
// @Summary Get product by ID
// @Description Return a specific product by its identifier
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 200 {object} entities.Product
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewGetProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	getProductCommand := commands.NewGetProductCommand(handler.ProductRepository)

	product, err := getProductCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// List product godoc
// @Summary List products
// @Description Return all registered products
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "pagination limit"
// @Param sort query string false "pagination sorting"
// @Success 200 {array} entities.Product
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewListProductsInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	listProductsCommand := commands.NewListProductsCommand(handler.ProductRepository)

	products, err := listProductsCommand.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Update product godoc
// @Summary Update product
// @Description Update product information
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Param request body inputs.UpdateProductInput true "product request"
// @Success 200
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewUpdateProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete product godoc
// @Summary Delete product
// @Description Remove product data
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "product ID" Format(uuid)
// @Success 200
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	input, err := inputs.NewDeleteProductInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
