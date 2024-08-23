package Middlewares

import (
	"github.com/gofiber/fiber/v2"
	"goSample/Services"
	"goSample/Types/Http"
	"log"
)

func VerifyToken(ctx *fiber.Ctx) error {
	token := ctx.Get("token")
	if token == "" {
		return Http.CreateHttpError(fiber.StatusUnauthorized, "Token không hợp lệ.")
	}
	tokenData, errorJWT := Services.ValidateTokenJWT(token)

	if errorJWT != nil {
		return Http.CreateHttpError(fiber.StatusUnauthorized, errorJWT.Error())
	}

	ctx.Locals("tokenData", tokenData)

	claimsData := ctx.Locals("tokenData").(*Services.Claims)
	log.Print(claimsData.UserID)
	return ctx.Next()
}
