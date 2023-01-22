package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Get("/links", GetAllLinks)
	router.Post("/links", CreateLink)
	router.Get("/links/:id", GetLinkById)
	router.Put("/links/:id", UpdateLink)
	router.Delete("/links/:id", DeleteLink)

	router.Get("/r/:shortenedId", Redirect)

	router.Listen(":8000")
}
