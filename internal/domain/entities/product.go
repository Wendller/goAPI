package entities

import (
	"errors"
	"time"

	"github.com/Wendller/goexpert/apis/pkg/entity"
)

var (
	ErrInvalidID     = errors.New("invalid id")
	ErrInvalidPrice  = errors.New("invalid price")
	ErrRequiredID    = errors.New("required id")
	ErrRequiredName  = errors.New("required name")
	ErrRequiredPrice = errors.New("required price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrRequiredID
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrRequiredName
	}

	if p.Price == 0 {
		return ErrRequiredPrice
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}
