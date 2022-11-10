package cart

import (
	"ecommerce/codebase/entity"
	"ecommerce/codebase/entity/request"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/repository/cart"
	"ecommerce/codebase/repository/product"

	"github.com/jinzhu/copier"
)

type ICartUseCase interface{}

type CartUseCase struct {
	cartRepository    cart.ICartRepo
	productRepository product.IProductRepo
}

func NewCartUseCase(cartRepository cart.ICartRepo, productRepository product.IProductRepo) *CartUseCase {
	return &CartUseCase{cartRepository: cartRepository, productRepository: productRepository}
}

func (c CartUseCase) CreateCart(req request.CreateCartRequest) error {
	product, err := c.productRepository.GetProductbyID(req.ProductID)
	if err != nil {
		return err
	}
	cart := entity.Cart{
		Qty:       req.Qty,
		ProductID: req.ProductID,
		Product:   product,
	}

	if err := c.cartRepository.CreateCart(cart); err != nil {
		return err
	}
	return nil
}

func (c CartUseCase) GetList() ([]response.GetCartResponse, error) {
	carts, err := c.cartRepository.GetAll()
	if err != nil {
		return nil, err
	}
	cartsResponse := []response.GetCartResponse{}
	copier.Copy(&cartsResponse, &carts)
	for k, v := range carts {
		product, err := c.productRepository.GetProductbyID(v.ProductID)
		if err != nil {
			return []response.GetCartResponse{}, err
		}
		joinedproduct, err := c.productRepository.JoinWithCategory(product, v.ProductID)
		if err != nil {
			return []response.GetCartResponse{}, err
		}
		productResponse := response.GetProductDetailResponse{}
		copier.Copy(&productResponse, &joinedproduct)
		productResponse.Category = joinedproduct.Category.Type
		copier.Copy(&cartsResponse[k].Product, productResponse)
	}
	return cartsResponse, nil
}
