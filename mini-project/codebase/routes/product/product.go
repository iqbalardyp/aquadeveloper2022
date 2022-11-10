package product

import (
	productHandler "ecommerce/codebase/handler/product"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo, productHandler *productHandler.ProductHandler) {
	r := app.Group("/api/v1")

	r.GET("/products", productHandler.GetList)
	r.GET("/products/price", productHandler.GetListbyPrice)
	r.GET("/products/category", productHandler.GetListbyCategory)
	r.POST("/product", productHandler.CreateProduct)
	r.GET("/product", productHandler.GetProduct)
}
