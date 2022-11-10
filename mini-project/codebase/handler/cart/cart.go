package cart

import (
	"ecommerce/codebase/entity/request"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/usecase/cart"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase *cart.CartUseCase
}

func NewCartHandler(cart *cart.CartUseCase) *CartHandler {
	return &CartHandler{cartUseCase: cart}
}

func (c CartHandler) CreateCart(ctx echo.Context) error {
	cartRequest := request.CreateCartRequest{}
	if err := ctx.Bind(&cartRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := c.cartUseCase.CreateCart(cartRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to create cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "cart created succesfully",
		Data:    cartRequest,
	})
}

func (c CartHandler) GetList(ctx echo.Context) error {
	carts, err := c.cartUseCase.GetList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "get list carts failed",
			Data:    nil,
		})
	}

	if len(carts) < 0 {
		return ctx.JSON(http.StatusNoContent, response.BaseResponse{
			Code:    http.StatusNoContent,
			Message: "no cart found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "succesfully get all carts",
		Data:    carts,
	})
}
