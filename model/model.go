package model

import (
	"github.com/amirfakhrullah/go-bitly/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Link struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	RedirectUrl string `json:"redirect_url" gorm:"unique;not null"`
	ShortenedId string `json:"shortened_id" gorm:"unique;not null"`
	Clicked     uint64 `json:"clicked"`
}

func SetupDB() {
	dsn := env.DB_URL_STRING
	var dbErr error
	db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	dbErr = db.AutoMigrate(&Link{})
	if dbErr != nil {
		panic(dbErr)
	}
}
