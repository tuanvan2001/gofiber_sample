package routes

import (
	"github.com/gofiber/fiber/v2"
	"goSample/controller"
)

func AuthRoutes(app *fiber.App) {
	authRoutes := app.Group("/user")

	authRoutes.Post("/login", controller.CreateUser)
}
