package payment

import (
	paymentHandler "ecommerce/codebase/handler/payment"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo, paymentHandler *paymentHandler.PaymentHandler) {
	r := app.Group("/api/v1")

	r.POST("/payment", paymentHandler.CreatePayment)
}
