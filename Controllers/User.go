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
	// check body & validate body
	if err := ctx.BodyParser(&createUserDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if errValidate, hasErrorValidate := Middlewares.Validator.Validate(createUserDto, Requests.CreateUserMessage); hasErrorValidate {
		return Http.CreateHttpErrorValidate(errValidate)
	}
	//call service create user
	id, errCreateUser := Services.CreateUser(createUserDto)
	if errCreateUser != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, errCreateUser.Error())

	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": Messages.User["Created"],
		"data": fiber.Map{
			"userID": id,
		},
	})
}
