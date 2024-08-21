package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Middlewares"
	"goSample/Services"
	"goSample/Types/Http"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

func CreateUser(ctx *fiber.Ctx) error {
	var createUserDto *Requests.CreateUser

	if err := ctx.BodyParser(&createUserDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if errMsgs, hasErrors := Middlewares.Validator.Validate(createUserDto, Requests.CreateUserMessage); hasErrors {
		return Http.CreateHttpErrorValidate(errMsgs)
	}
	id, errCreateUser := Services.CreateUser(createUserDto)
	if errCreateUser != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, errCreateUser.Error())

	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": UserMessages.Vi["Created"],
		"data": fiber.Map{
			"userID": id,
		},
	})
}
