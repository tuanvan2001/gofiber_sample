package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Middlewares"
	"goSample/Types/Http"
	"goSample/Types/Requests"
)

func CreateUser(ctx *fiber.Ctx) error {
	var createUserDto *Requests.CreateUser
	if err := ctx.BodyParser(&createUserDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if errMsgs, hasErrors := Middlewares.Validator.Validate(createUserDto); hasErrors {
		return Http.CreateHttpError(fiber.StatusBadRequest, "Validation failed", errMsgs)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"token": "abcxyz",
		},
	})
}
