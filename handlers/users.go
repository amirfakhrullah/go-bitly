package handlers

import (
	"time"

	"github.com/amirfakhrullah/go-bitly/env"
	"github.com/amirfakhrullah/go-bitly/helpers"
	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	type LoginPayload struct {
		Email    string `json:"email" validate:"required,email,min=4,max=50"`
		Password string `json:"password" validate:"required,min=4,max=20"`
	}

	var body LoginPayload
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}
	if err := helpers.ValidatePayload(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	user, err := services.FindUserByEmail(body.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if ok := helpers.CheckPasswordHash(body.Password, user.HashedPassword); !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// create jwt token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	ctx.Append("Authorization", "Bearer "+t)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}

func Signup(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	type SignupPayload struct {
		Name     string `json:"name" validate:"required,min=4,max=50"`
		Email    string `json:"email" validate:"required,email,min=4,max=50"`
		Password string `json:"password" validate:"required,min=8,max=20"`
	}

	var body SignupPayload
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}

	if err := helpers.ValidatePayload(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	isExists, err := services.IsUserExists(body.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if isExists {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	hashedPassword, err := helpers.HashPassword(body.Password)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if err := services.CreateUser(model.User{
		Name:           body.Name,
		Email:          body.Email,
		HashedPassword: hashedPassword,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating a new user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user successfully created",
	})
}

func Logout(ctx *fiber.Ctx) error {
	ctx.Append("Authorization", "")
	return ctx.SendStatus(fiber.StatusOK)
}

func GetUserInfo(ctx *fiber.Ctx) error {
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user, err := services.GetUserInfo(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}
