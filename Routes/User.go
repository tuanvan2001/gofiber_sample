package Routes

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Controllers"
)

func UserRoutes(app *fiber.App) {
	authRoutes := app.Group("/user")

	authRoutes.Post("/login", Controllers.CreateUser)
}
