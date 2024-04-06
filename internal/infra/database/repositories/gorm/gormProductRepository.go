package gorm_repositories

import (
	"errors"

	"github.com/Wendller/goexpert/apis/internal/domain/entities"
	"gorm.io/gorm"
)

type GORMProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *GORMProductRepository {
	return &GORMProductRepository{DB: db}
}

func (p *GORMProductRepository) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *GORMProductRepository) FindByID(id string) (*entities.Product, error) {
	var product entities.Product

	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *GORMProductRepository) FindMany(page, limit int, sort string) ([]entities.Product, error) {
	var products []entities.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		err = errors.New("invalid sort")
		return nil, err
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *GORMProductRepository) Update(product *entities.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *GORMProductRepository) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
