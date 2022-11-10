package category

import (
	categoryHandler "ecommerce/codebase/handler/category"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo, categoryHandler *categoryHandler.CategoryHandler) {
	r := app.Group("/api/v1")

	r.GET("/categories", categoryHandler.GetList)
	r.POST("/category", categoryHandler.CreateCategory)
	r.GET("/category", categoryHandler.GetCategory)
}
