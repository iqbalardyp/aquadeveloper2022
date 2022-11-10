package category

import (
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/usecase/category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUseCase *category.CategoryUseCase
}

func NewCategoryHandler(category *category.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: category}
}

func (c CategoryHandler) CreateCategory(ctx echo.Context) error {
	categoryRequest := response.GetCategoryResponse{}
	if err := ctx.Bind(&categoryRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := c.categoryUseCase.CreateCategory(categoryRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to create category",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "category created succesfully",
		Data:    categoryRequest,
	})
}

func (c CategoryHandler) GetList(ctx echo.Context) error {
	categories, err := c.categoryUseCase.GetList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list categories failed",
			Data:    nil,
		})
	}

	if len(categories) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no category found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "succesfully get all categories",
		Data:    categories,
	})
}

func (c CategoryHandler) GetCategory(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.FormValue("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "no id found",
			Data:    nil,
		})
	}
	category, err := c.categoryUseCase.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get category failed",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get category",
		Data:    category,
	})
}
