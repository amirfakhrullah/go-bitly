package main

import (
	"os"

	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("DB_URL_STRING")
	model.SetupDB(dsn)
	server.SetupAndListen()
}
