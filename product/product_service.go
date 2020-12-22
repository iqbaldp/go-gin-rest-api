package product

type ProductService struct {
	ProductRepository ProductRepository
}

func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) findAll() []Product {
	return p.ProductRepository.findAll()
}

func (p *ProductService) findByID(id uint) Product {
	return p.ProductRepository.findByID(id)
}

func (p *ProductService) Save(product Product) Product {
	p.ProductRepository.save(product)

	return product
}

func (p *ProductService) Delete(product Product) {
	p.ProductRepository.delete(product)
}
