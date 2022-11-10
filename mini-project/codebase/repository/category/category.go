package category

import (
	"ecommerce/codebase/entity"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

type ICategoryRepo interface {
	Create(category entity.Category) error
	GetAll() ([]entity.Category, error)
	GetCategorybyID(id int64) (entity.Category, error)
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (c CategoryRepo) Create(category entity.Category) error {
	if err := c.db.Debug().Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepo) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	if err := c.db.Debug().Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c CategoryRepo) GetCategorybyID(id int64) (entity.Category, error) {
	var category entity.Category
	if err := c.db.Debug().Find(&category, id).Error; err != nil {
		return entity.Category{}, err
	}
	return category, nil
}
