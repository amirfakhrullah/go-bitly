package middlewares

import (
	jwtware "github.com/gofiber/jwt/v3"
)

var AuthMiddleware = jwtware.New(jwtware.Config{
	SigningKey: []byte("secret"),
})
