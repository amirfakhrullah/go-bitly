package main

import (
	"log"
	"os"

	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL_STRING")
	model.SetupDB(dsn)
	server.SetupAndListen()
}
