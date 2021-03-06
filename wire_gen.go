// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"rest-api/product"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:

func initProductAPI(db *gorm.DB) product.ProductAPI {
	productRepository := product.ProvideProductRepository(db)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductApi(productService)
	return productAPI
}
