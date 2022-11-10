package cart

import (
	"ecommerce/codebase/entity"

	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

type ICartRepo interface {
	CreateCart(cart entity.Cart) error
	GetAll() ([]entity.Cart, error)
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (c CartRepo) CreateCart(cart entity.Cart) error {
	if err := c.db.Debug().Create(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (c CartRepo) GetAll() ([]entity.Cart, error) {
	var cartList []entity.Cart
	if err := c.db.Debug().Joins("Product").Find(&cartList).Error; err != nil {
		return nil, err
	}
	return cartList, nil
}
