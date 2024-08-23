package Routes

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Controllers"
	"goSample/Middlewares"
)

func UserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user")

	userRoutes.Post("/", Middlewares.VerifyToken, Controllers.CreateUser)
}
