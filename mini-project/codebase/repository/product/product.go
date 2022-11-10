package product

import (
	"ecommerce/codebase/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

type IProductRepo interface {
	Create(product entity.Product) error
	GetAll() ([]entity.Product, error)
	GetProductbyID(id int64) (entity.Product, error)
	JoinWithCategory(product entity.Product, id int64) (entity.Product, error)
	JoinWithCategories(product []entity.Product) ([]entity.Product, error)
	FilterbyPrice(priceMin int64, priceMax int64) ([]entity.Product, error)
	FilterbyCategory(category string) ([]entity.Product, error)
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p ProductRepo) Create(product entity.Product) error {
	if err := p.db.Debug().Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepo) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Debug().Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductRepo) GetProductbyID(id int64) (entity.Product, error) {
	var product entity.Product
	if err := p.db.Debug().Find(&product, &id).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p ProductRepo) DeleteProductbyID(id int64) error {
	var product entity.Product
	if err := p.db.Debug().Delete(&product, &id).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepo) UpdateProduct(product entity.Product) (entity.Product, error) {
	if err := p.db.Debug().Save(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p ProductRepo) JoinWithCategory(product entity.Product, id int64) (entity.Product, error) {
	if err := p.db.Joins("Category").Find(&product, &id).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p ProductRepo) JoinWithCategories(product []entity.Product) ([]entity.Product, error) {
	if err := p.db.Joins("Category").Find(&product).Error; err != nil {
		return []entity.Product{}, err
	}
	return product, nil
}

func (p ProductRepo) FilterbyPrice(priceMin int64, priceMax int64) ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Joins("Category").
		Where("products.price >= ? and products.price <= ?", priceMin, priceMax).
		Find(&products).Error; err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}

func (p ProductRepo) FilterbyCategory(category string) ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Joins("Category").
		Where("type = ?", category).
		Find(&products).Error; err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}
