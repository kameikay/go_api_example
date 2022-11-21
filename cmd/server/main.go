package main

import (
	"net/http"

	"github.com/kameikay/api_example/configs"
	"github.com/kameikay/api_example/infra/database"
	"github.com/kameikay/api_example/infra/webserver/handlers"
	"github.com/kameikay/api_example/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Product{}, &entities.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
