package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserId (ctx *fiber.Ctx) uint {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(uint)
	if id == 0 {
		ctx.SendStatus(fiber.StatusUnauthorized)
	}
	return id
}