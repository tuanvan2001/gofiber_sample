package Middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func VerifyToken(ctx *fiber.Ctx) error {
	token := ctx.Get("token")
	log.Println(token)
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid token",
		})
	}
	return ctx.Next()
}
