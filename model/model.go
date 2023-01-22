package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Link struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	RedirectUrl string `json:"redirectUrl" gorm:"unique;not null"`
	ShortenedId string `json:"shortenedId" gorm:"unique;not null"`
	Clicked     uint64 `json:"clicked"`
}

func SetupDB(dsn string) {
	if dsn == "" {
		panic("database url not found")
	}
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
