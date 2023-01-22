package main

import (
	"log"

	"github.com/amirfakhrullah/go-bitly/db"
	"github.com/amirfakhrullah/go-bitly/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.ConnectDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetupLinkRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
