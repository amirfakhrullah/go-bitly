package middlewares

import (
	"github.com/amirfakhrullah/go-bitly/env"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey: []byte(env.JWT_SECRET),
	ErrorHandler: func(ctx *fiber.Ctx, _ error) error {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	},
})
