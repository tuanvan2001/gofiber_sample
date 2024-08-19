package main

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Configs"
	"goSample/Routes"
	Http "goSample/Types/Http"
	"log"
)

func init() {
	Configs.LoadENV()
	Configs.ConnectMySQL()
	Configs.ConnectRedis()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			type ErrorResponse struct {
				Message string      `json:"message"`
				Code    int         `json:"code"`
				Detail  interface{} `json:"detail,omitempty"`
			}

			if e, ok := err.(*Http.HttpError); ok {
				customError := ErrorResponse{
					Message: e.Message,
					Code:    e.Code,
					Detail:  e.Detail,
				}
				return ctx.Status(e.Code).JSON(customError)
			}
			if e, ok := err.(*fiber.Error); ok {
				customError := ErrorResponse{
					Message: e.Message,
					Code:    e.Code,
				}
				return ctx.Status(e.Code).JSON(customError)
			}
			customError := ErrorResponse{
				Message: "Internal Server Error.",
				Code:    fiber.StatusInternalServerError,
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(customError)
		},
	})

	app.Get("/healthCheck", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Ok",
		})
	})
	Routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
