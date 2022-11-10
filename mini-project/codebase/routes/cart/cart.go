package cart

import (
	cartHandler "ecommerce/codebase/handler/cart"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo, cartHandler *cartHandler.CartHandler) {
	r := app.Group("/api/v1")

	r.GET("/cart", cartHandler.GetList)
	r.POST("/cart", cartHandler.CreateCart)
}
