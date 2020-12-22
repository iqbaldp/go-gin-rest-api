//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"rest-api/product"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductApi)

	return product.ProductAPI{}
}
