package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Middlewares"
	"goSample/Services"
	"goSample/Types/Http"
	"goSample/Types/Messages"
	"goSample/Types/Requests"
	"log"
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

	user, err := Services.Login(loginDto)
	if err != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, err.Error())
	}

	claims, errJWT := Services.ValidateTokenJWT(user.Token)
	if errJWT != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, err.Error())
	}
	log.Println(claims.Username)
	log.Println(claims.UserID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": Messages.Login["Success"],
		"data":    user,
	})
}

func Register(ctx *fiber.Ctx) error {
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

	user, err := Services.Login(loginDto)
	if err != nil {
		return Http.CreateHttpError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": Messages.Login["Success"],
		"data":    user,
	})
}
