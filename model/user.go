package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	HashedPassword string `json:"-" gorm:"unique;not null"`
	Links          []Link `json:"links" gorm:"foreignKey:UserID"`
}
