package product

import (
	"ecommerce/codebase/entity/request"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/usecase/product"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase *product.ProductUseCase
}

func NewProductHandler(product *product.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: product}
}

func (p ProductHandler) CreateProduct(ctx echo.Context) error {
	productRequest := request.CreateProductDetailRequest{}
	if err := ctx.Bind(&productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	price, _ := strconv.ParseInt(ctx.FormValue("price"), 10, 64)
	category, _ := strconv.ParseInt(ctx.FormValue("categoryID"), 10, 64)

	// Source
	formImg, err := ctx.FormFile("figure")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to read form file",
			Data:    nil,
		})
	}
	openImg, err := formImg.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to read file",
			Data:    nil,
		})
	}
	defer openImg.Close()

	readImg, _ := ioutil.ReadAll(openImg)

	productRequest.Name = ctx.FormValue("name")
	productRequest.Price = price
	productRequest.CategoryID = category
	productRequest.Desc = ctx.FormValue("desc")
	productRequest.Figure = readImg
	if err := p.productUseCase.CreateProduct(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to create product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "product created succesfully",
		Data:    productRequest,
	})
}

func (p ProductHandler) GetList(ctx echo.Context) error {
	products, err := p.productUseCase.GetList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list products failed",
			Data:    nil,
		})
	}

	if len(products) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no product found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "succesfully get all products",
		Data:    products,
	})
}

func (p ProductHandler) GetListbyPrice(ctx echo.Context) error {
	priceMin, _ := strconv.ParseInt(ctx.FormValue("priceMin"), 10, 64)
	priceMax, _ := strconv.ParseInt(ctx.FormValue("priceMax"), 10, 64)
	if priceMin > priceMax {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	products, err := p.productUseCase.GetByPriceRange(priceMin, priceMax)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list products failed",
			Data:    nil,
		})
	}
	if len(products) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no product found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "succesfully get all products",
		Data:    products,
	})
}

func (p ProductHandler) GetListbyCategory(ctx echo.Context) error {
	products, err := p.productUseCase.GetByCategory(ctx.FormValue("category"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list products failed",
			Data:    nil,
		})
	}
	if len(products) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no product found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "succesfully get all products",
		Data:    products,
	})
}

func (p ProductHandler) GetProduct(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.FormValue("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "id not found",
			Data:    nil,
		})
	}
	product, err := p.productUseCase.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get product failed",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "successfully get product",
		Data:    product,
	})
}
