package category

import (
	"ecommerce/codebase/entity"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/repository/category"

	"github.com/jinzhu/copier"
)

type ICategoryUseCase interface{}

type CategoryUseCase struct {
	categoryRepository category.ICategoryRepo
}

func NewCategoryUseCase(categoryRepository category.ICategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{categoryRepository: categoryRepository}
}

func (c CategoryUseCase) CreateCategory(req response.GetCategoryResponse) error {
	categories := entity.Category{}
	copier.Copy(&categories, &req)

	err := c.categoryRepository.Create(categories)
	if err != nil {
		return err
	}
	return nil
}

func (c CategoryUseCase) GetList() ([]response.GetCategoryResponse, error) {
	categories, err := c.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	categoryResponse := []response.GetCategoryResponse{}
	copier.Copy(&categoryResponse, &categories)
	return categoryResponse, nil
}

func (c CategoryUseCase) GetByID(id int64) (response.GetCategoryResponse, error) {
	category, err := c.categoryRepository.GetCategorybyID(id)
	if err != nil {
		return response.GetCategoryResponse{}, err
	}
	categoryResponse := response.GetCategoryResponse{}
	copier.Copy(&categoryResponse, &category)
	return categoryResponse, nil
}
