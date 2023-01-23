package services

import (
	"errors"

	"github.com/amirfakhrullah/go-bitly/db"
	"github.com/amirfakhrullah/go-bitly/model"
)

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := db.DB.Find(&user, "email = ?", email).Error; err != nil {
		return model.User{}, err
	}
	if user.ID == 0 {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func IsUserExists(email string) (bool, error) {
	var user model.User
	if err := db.DB.Find(&user, "email = ?", email).Error; err != nil {
		return false, err
	}
	return user.ID != 0, nil
}

func CreateUser(user model.User) error {
	tx := db.DB.Create(&user)
	return tx.Error
}