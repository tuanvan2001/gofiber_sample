package Controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"goSample/Configs"
	"goSample/Models"
	"time"
)

type CreateUserRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Age         uint8  `json:"age" validate:"required,min=1"`
	Birthday    string `json:"birthday" validate:"required,datetime=2006-01-02"`
}

type UserController struct {
}

var validate = validator.New()

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {

	var body CreateUserRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	if validationErr := validate.Struct(&body); validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation error",
			"errors":  validationErr.Error(),
		})
	}
	var existingUserByEmail, existingUserByUsername, existingUserByPhone Models.User
	Configs.MySQL.Where("email = ?", body.Email).First(&existingUserByEmail)

	Configs.MySQL.Where("username = ?", body.Username).First(&existingUserByUsername)

	Configs.MySQL.Where("phone_number = ?", body.PhoneNumber).First(&existingUserByPhone)

	birthday, _ := time.Parse("2006-01-02", body.Birthday)

	user := Models.User{
		Username:    body.Username,
		Password:    body.Password,
		FullName:    body.FullName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		Age:         body.Age,
		Birthday:    birthday,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save user to database
	if result := Configs.MySQL.Create(&user); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating user",
			"error":   result.Error.Error(),
		})
	}

	// Create an event for the new user creation
	event := Models.Event{
		EventName: "Tạo user",
		UserID:    int(user.ID), // Assuming UserID is the ID of the newly created user
		CreatedAt: time.Now(),
	}

	if result := Configs.MySQL.Create(&event); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating event",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"user":  user,
			"event": event,
		},
	})
}

func (uc *UserController) ListUsers(ctx *fiber.Ctx) error {
	var users []Models.User

	if result := Configs.MySQL.Find(&users); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error retrieving users",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}
