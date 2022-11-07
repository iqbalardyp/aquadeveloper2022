package main

import (
	"demo-clean-arch/config"
	"demo-clean-arch/handler"
	"demo-clean-arch/repository"
	"demo-clean-arch/routes"
	"demo-clean-arch/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	// var user1 = entity.User{
	// 	ID:      1,
	// 	Name:    "iqbal",
	// 	Email:   "iqbal.ardyp@gmail.com",
	// 	Phone:   "081232824347",
	// 	Address: "cisitu lama",
	// }

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(app, userHandler)
	app.Listen(":5555")
}
