package Routes

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Controllers"
)

func AuthRoutes(app *fiber.App) {
	authRoutes := app.Group("/users")
	userController := &Controllers.UserController{}

	authRoutes.Get("/", userController.ListUsers)
	authRoutes.Post("/", userController.CreateUser)
}
