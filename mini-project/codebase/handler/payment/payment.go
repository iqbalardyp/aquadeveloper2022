package payment

import (
	"ecommerce/codebase/entity"
	"ecommerce/codebase/entity/response"
	"ecommerce/codebase/usecase/payment"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentUseCase *payment.PaymentUseCase
}

func NewPaymentHandler(payment *payment.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{paymentUseCase: payment}
}

func (p PaymentHandler) CreatePayment(ctx echo.Context) error {
	paymentRequest := entity.Payment{}
	if err := ctx.Bind(&paymentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	formPayment, err := ctx.FormFile("paymentSlip")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to read form file",
			Data:    nil,
		})
	}
	openPayment, err := formPayment.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to read file",
			Data:    nil,
		})
	}
	defer openPayment.Close()

	readPayment, _ := ioutil.ReadAll(openPayment)

	paymentRequest.PaymentSlip = readPayment
	paymentResponse, err := p.paymentUseCase.CreatePayment(paymentRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "failed to create payment",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "payment created succesfully",
		Data:    paymentResponse,
	})
}
