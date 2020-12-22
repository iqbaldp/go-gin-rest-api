package product

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProvideProductApi(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(context *gin.Context) {
	products := p.ProductService.findAll()

	context.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductAPI) FindByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	product := p.ProductService.findByID(uint(id))

	context.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductAPI) Create(context *gin.Context) {
	var productDTO ProductDTO
	err := context.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		context.Status(http.StatusBadRequest)
		return
	}

	createProduct := p.ProductService.Save(ToProduct(productDTO))
	context.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createProduct)})
}

func (p *ProductAPI) Update(context *gin.Context) {
	var productDTO ProductDTO
	err := context.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		context.Status(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(context.Param("id"))
	product := p.ProductService.findByID(uint(id))

	if product == (Product{}) {
		context.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	context.Status(http.StatusOK)

	createProduct := p.ProductService.Save(ToProduct(productDTO))
	context.JSON(http.StatusOK, gin.H{"product": ToProductDTO(createProduct)})
}

func (p *ProductAPI) Delete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	product := p.ProductService.findByID(uint(id))

	if product == (Product{}) {
		context.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	context.Status(http.StatusOK)
}
