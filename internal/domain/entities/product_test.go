package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductCreation(t *testing.T) {
	t.Run("Valid Product", func(t *testing.T) {
		product, err := NewProduct("Fighter Z", float64(float64(30)))

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.NotEmpty(t, product.ID)
		assert.Equal(t, "Fighter Z", product.Name)
		assert.Equal(t, float64(30), product.Price)
	})

	t.Run("Product Without Name", func(t *testing.T) {
		product, err := NewProduct("", float64(30))

		assert.Nil(t, product)
		assert.Equal(t, ErrRequiredName, err)
	})

	t.Run("Product Without Price", func(t *testing.T) {
		product, err := NewProduct("Product 1", 0.0)

		assert.Nil(t, product)
		assert.Equal(t, ErrRequiredPrice, err)
	})

	t.Run("Product With Invalid Price", func(t *testing.T) {
		product, err := NewProduct("Product 1", -10.0)

		assert.Nil(t, product)
		assert.Equal(t, ErrInvalidPrice, err)
	})

	t.Run("Product Validation", func(t *testing.T) {
		product, err := NewProduct("Product 1", 10.0)

		assert.Nil(t, err)
		assert.Nil(t, product.Validate())
	})
}
