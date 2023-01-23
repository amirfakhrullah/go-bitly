package routes

import (
	"github.com/amirfakhrullah/go-bitly/handlers"
	"github.com/amirfakhrullah/go-bitly/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupLinkRoutes(app *fiber.App) {
	r := app.Group("/links", middlewares.JwtMiddleware)

	r.Get("/", handlers.GetAllLinks)
	r.Post("/", handlers.CreateLink)
	r.Get("/:id", handlers.GetLinkById)
	r.Put("/:id", handlers.UpdateLink)
	r.Delete("/:id", handlers.DeleteLink)
}
