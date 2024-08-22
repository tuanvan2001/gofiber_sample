package Routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	UserRoutes(app)
	AuthRoutes(app)
}
