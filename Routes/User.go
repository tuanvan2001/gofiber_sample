package Routes

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Controllers"
)

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user")

	userRoutes.Post("/", Controllers.CreateUser)
}
