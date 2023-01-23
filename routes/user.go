package routes

import (
	"github.com/amirfakhrullah/go-bitly/handlers"
	"github.com/amirfakhrullah/go-bitly/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	r := app.Group("/user", middlewares.JwtMiddleware)

	r.Get("/", handlers.GetUserInfo)
}
