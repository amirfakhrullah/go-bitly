package db

import (
	"log"
	"os"

	"github.com/amirfakhrullah/go-bitly/env"
	"github.com/amirfakhrullah/go-bitly/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := env.DB_URL_STRING
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	DB.AutoMigrate(&model.Link{})
}
