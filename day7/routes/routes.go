package routes

import (
	"demo-clean-arch/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Get("/users", userHandler.GetList)
	r.Post("/user", userHandler.CreateUser)
	r.Get("/user/:id", userHandler.GetUser)
	r.Put("/user/:id", userHandler.UpdateUserByID)
	r.Delete("/user/:id", userHandler.DeleteUser)
}
