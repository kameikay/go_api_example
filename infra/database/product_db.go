package database

import (
	"github.com/kameikay/api_example/internal/entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {
	var products []entities.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

func (p *Product) FindByID(id string) (*entities.Product, error) {
	var product entities.Product
	err := p.DB.First(&product, "id = ?", id).Error

	return &product, err
}

func (p *Product) Update(product *entities.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
