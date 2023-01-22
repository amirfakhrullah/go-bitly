package routes

import (
	"github.com/amirfakhrullah/go-bitly/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupLinkRoutes(r *fiber.App) {
	r.Get("/links", handlers.GetAllLinks)
	r.Post("/links", handlers.CreateLink)
	r.Get("/links/:id", handlers.GetLinkById)
	r.Put("/links/:id", handlers.UpdateLink)
	r.Delete("/links/:id", handlers.DeleteLink)

	r.Get("/r/:shortenedId", handlers.Redirect)
}
