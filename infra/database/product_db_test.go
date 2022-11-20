package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/kameikay/api_example/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}

	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)

	var productFound entities.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.Product{})

	for i := 1; i < 24; i++ {
		product, err := entities.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}
	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("Product 1", 123.0)
	assert.Nil(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}

	db.AutoMigrate(&entities.Product{})

	product, err := entities.NewProduct("Product 1", 123.0)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.Nil(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error()
	}

	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("Product 1", 123.0)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
