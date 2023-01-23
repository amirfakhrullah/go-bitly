package routes

import (
	"github.com/amirfakhrullah/go-bitly/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRedirectRoutes(app *fiber.App) {
	r := app.Group("/r")

	r.Get("/:shortenedId", handlers.Redirect)
}
