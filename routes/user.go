package routes

import (
	"github.com/amirfakhrullah/go-bitly/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	r := app.Group("/user")

	r.Post("/login", handlers.Login)
	r.Post("/signup", handlers.Signup)
	r.Post("/logout", handlers.Logout)
}