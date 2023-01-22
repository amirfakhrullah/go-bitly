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
	router.Get("/links/:id", GetLinkById)
	router.Post("/links", CreateLink)

	router.Listen(":8000")
}
