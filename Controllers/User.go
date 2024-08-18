package Controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"token": "abcxyz",
		},
	})
}
