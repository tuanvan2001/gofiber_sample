package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Middlewares"
	"goSample/Services"
	"goSample/Types/Http"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
)

func Login(ctx *fiber.Ctx) error {
	var loginDto *Requests.LoginDto

	if err := ctx.BodyParser(&loginDto); err != nil {
		return Http.CreateHttpError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}
	if errMsgs, hasErrors := Middlewares.Validator.Validate(loginDto, Requests.LoginMessage); hasErrors {
		return Http.CreateHttpErrorValidate(errMsgs)
	}

	user, _ := Services.Login(loginDto)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": Messages.Login["Success"],
		"data":    user,
	})
}
