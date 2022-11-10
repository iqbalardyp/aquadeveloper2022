package product

import (
	"ecommerce/codebase/entity"
	"ecommerce/codebase/entity/request"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/repository/product"

	"github.com/jinzhu/copier"
)

type IProductUseCase interface{}

type ProductUseCase struct {
	productRepository product.IProductRepo
}

func NewProductUseCase(productRepository product.IProductRepo) *ProductUseCase {
	return &ProductUseCase{productRepository: productRepository}
}

func (p ProductUseCase) CreateProduct(req request.CreateProductDetailRequest) error {
	var product entity.Product
	copier.Copy(&product, &req)

	err := p.productRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductUseCase) GetList() ([]response.BaseProductResponse, error) {
	products, err := p.productRepository.GetAll()
	if err != nil {
		return nil, err
	}
	joinedproducts, err := p.productRepository.JoinWithCategories(products)
	if err != nil {
		return []response.BaseProductResponse{}, err
	}
	productsRespons := []response.BaseProductResponse{}
	copier.Copy(&productsRespons, &joinedproducts)
	for i, v := range joinedproducts {
		productsRespons[i].Category = v.Category.Type
	}
	return productsRespons, nil
}

func (p ProductUseCase) GetByID(id int64) (response.GetProductDetailResponse, error) {
	product, err := p.productRepository.GetProductbyID(id)
	if err != nil {
		return response.GetProductDetailResponse{}, err
	}
	joinedproduct, err := p.productRepository.JoinWithCategory(product, id)
	if err != nil {
		return response.GetProductDetailResponse{}, err
	}
	productResponse := response.GetProductDetailResponse{}
	copier.Copy(&productResponse, &joinedproduct)
	productResponse.Category = joinedproduct.Category.Type
	return productResponse, nil
}

func (p ProductUseCase) GetByPriceRange(priceMin int64, priceMax int64) ([]response.BaseProductResponse, error) {
	products, err := p.productRepository.FilterbyPrice(priceMin, priceMax)
	if err != nil {
		return []response.BaseProductResponse{}, err
	}
	productsResponse := []response.BaseProductResponse{}
	copier.Copy(&productsResponse, &products)
	for i, v := range products {
		productsResponse[i].Category = v.Category.Type
	}
	return productsResponse, nil
}

func (p ProductUseCase) GetByCategory(category string) ([]response.BaseProductResponse, error) {
	products, err := p.productRepository.FilterbyCategory(category)
	if err != nil {
		return []response.BaseProductResponse{}, err
	}
	productsResponse := []response.BaseProductResponse{}
	copier.Copy(&productsResponse, &products)
	for i, v := range products {
		productsResponse[i].Category = v.Category.Type
	}
	return productsResponse, nil
}
