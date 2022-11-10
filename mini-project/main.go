package main

import (
	"ecommerce/codebase/config"
	categoryHandler "ecommerce/codebase/handler/category"
	categoryRepo "ecommerce/codebase/repository/category"
	categoryRoutes "ecommerce/codebase/routes/category"
	categoryUseCase "ecommerce/codebase/usecase/category"

	productHandler "ecommerce/codebase/handler/product"
	productRepo "ecommerce/codebase/repository/product"
	productRoutes "ecommerce/codebase/routes/product"
	productUseCase "ecommerce/codebase/usecase/product"

	cartHandler "ecommerce/codebase/handler/cart"
	cartRepo "ecommerce/codebase/repository/cart"
	cartRoutes "ecommerce/codebase/routes/cart"
	cartUseCase "ecommerce/codebase/usecase/cart"

	paymentHandler "ecommerce/codebase/handler/payment"
	paymentRepo "ecommerce/codebase/repository/payment"
	paymentRoutes "ecommerce/codebase/routes/payment"
	paymentUseCase "ecommerce/codebase/usecase/payment"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := echo.New()
	app.Use(middleware.Logger())

	// Category
	categoryRepository := categoryRepo.NewCategoryRepo(config.DB)
	categoryUsecase := categoryUseCase.NewCategoryUseCase(categoryRepository)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryUsecase)

	categoryRoutes.Routes(app, categoryHandler)

	// Products
	productRepository := productRepo.NewProductRepo(config.DB)
	productUsecase := productUseCase.NewProductUseCase(productRepository)
	productHandler := productHandler.NewProductHandler(productUsecase)

	productRoutes.Routes(app, productHandler)

	// Cart
	cartRepository := cartRepo.NewCartRepo(config.DB)
	cartUsecase := cartUseCase.NewCartUseCase(cartRepository, productRepository)
	cartHandler := cartHandler.NewCartHandler(cartUsecase)

	cartRoutes.Routes(app, cartHandler)

	// Payment
	paymentRepository := paymentRepo.NewPaymentRepo(config.DB)
	paymentUsecase := paymentUseCase.NewPaymentUseCase(paymentRepository)
	paymentHandler := paymentHandler.NewPaymentHandler(paymentUsecase)

	paymentRoutes.Routes(app, paymentHandler)
	app.Logger.Fatal(app.Start(":1323"))
}
