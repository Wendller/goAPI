package database

import (
	"fmt"
	"os"
	"math/rand"
	"testing"

	"github.com/Wendller/goexpert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
    var err error
    DB, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    DB.AutoMigrate(&entity.Product{})

    code := m.Run()

    DB.Exec("DELETE FROM products")

    DBConn, err := DB.DB()
    if err != nil {
        panic("Failed to get database connection")
    }
    DBConn.Close()

    os.Exit(code)
}

func TestCreate(t *testing.T) {
	t.Run("Create new product", func(t *testing.T) {
		product, _ := entity.NewProduct("Product 1", 1500)
		productDB := NewProduct(DB)

		err := productDB.Create(product)
		assert.Nil(t, err)

		var createdProduct entity.Product
		err = DB.First(&createdProduct, "id = ?", product.ID).Error

		assert.Nil(t, err)
		assert.Equal(t, createdProduct.ID, product.ID)
	})
}

func TestFindMany(t *testing.T) {
	t.Run("Find many products ordered asc", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		for i := 1; i < 25; i++ {
			product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			DB.Create(product)
		}

		productDB := NewProduct(DB)
		products, err := productDB.FindMany(1, 10, "asc")

		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 10", products[9].Name)

		products, err = productDB.FindMany(2, 10, "asc")

		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 11", products[0].Name)
		assert.Equal(t, "Product 20", products[9].Name)

		products, err = productDB.FindMany(3, 10, "asc")

		assert.NoError(t, err)
		assert.Len(t, products, 4)
		assert.Equal(t, "Product 21", products[0].Name)
		assert.Equal(t, "Product 24", products[3].Name)		
	})

	t.Run("Find many products ordered desc", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		for i := 1; i < 25; i++ {
			product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			DB.Create(product)
		}

		productDB := NewProduct(DB)
		products, err := productDB.FindMany(1, 10, "desc")

		assert.NoError(t, err)
		assert.Len(t, products, 10)
		assert.Equal(t, "Product 24", products[0].Name)
		assert.Equal(t, "Product 15", products[9].Name)		
	})

	t.Run("Find many products with empty sort", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		for i := 1; i < 6; i++ {
			product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			DB.Create(product)
		}

		productDB := NewProduct(DB)
		products, err := productDB.FindMany(1, 5, "")

		assert.NoError(t, err)
		assert.Len(t, products, 5)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 5", products[4].Name)		
	})

	t.Run("Find many products when invalid sort", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		for i := 1; i < 6; i++ {
			product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			DB.Create(product)
		}

		productDB := NewProduct(DB)
		_, err = productDB.FindMany(1, 5, "invalid")

		assert.ErrorContains(t, err, "invalid sort")
	})

	t.Run("Find many products when page and limit are 0", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		for i := 1; i < 6; i++ {
			product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
			DB.Create(product)
		}

		productDB := NewProduct(DB)
		products, err := productDB.FindMany(0, 0, "")

		assert.NoError(t, err)
		assert.Len(t, products, 5)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, "Product 5", products[4].Name)		
	})
}

func TestFinDByID(t *testing.T) {
	t.Run("Find when product exists", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		product, _ := entity.NewProduct("Product 1", rand.Float64()*100)
		product_2, _ := entity.NewProduct("Product 2", rand.Float64()*100)
		DB.Create(product)
		DB.Create(product_2)
		

		productDB := NewProduct(DB)
		targetProduct, err := productDB.FindByID(product.ID.String())

		assert.NoError(t, err)
		assert.Equal(t, targetProduct.ID, product.ID)
	})

	t.Run("Find when product doens't exists", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		productDB := NewProduct(DB)
		product, err := productDB.FindByID("1")

		assert.Nil(t, product)
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update product", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		product, _ := entity.NewProduct("Product 1", rand.Float64()*100)
		DB.Create(product)

		product.Name = "Product 2"
		
		productDB := NewProduct(DB)
		err = productDB.Update(product)

		assert.NoError(t, err)
		
		updatedProduct, err := productDB.FindByID(product.ID.String())

		assert.NoError(t, err)
		assert.Equal(t, updatedProduct.Name, "Product 2")
	})

	t.Run("Fail when product doens't exists", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		product, _ := entity.NewProduct("Product 1", rand.Float64()*100)
		product.Name = "Product 2"
		
		productDB := NewProduct(DB)
		err = productDB.Update(product)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete product", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		product, _ := entity.NewProduct("Product 1", rand.Float64()*100)
		DB.Create(product)

		productDB := NewProduct(DB)
		err = productDB.Delete(product.ID.String())

		assert.NoError(t, err)
		
		targetProduct, err := productDB.FindByID(product.ID.String())
		
		assert.Nil(t, targetProduct)
		assert.Error(t, err)
	})

	t.Run("Fail when product doens't exists", func(t *testing.T) {
		DB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}

		DB.AutoMigrate(&entity.Product{})

		product, _ := entity.NewProduct("Product 1", rand.Float64()*100)
		
		productDB := NewProduct(DB)
		err = productDB.Delete(product.ID.String())

		assert.Error(t, err)
	})
}