package helpers

import (
	"errors"

	"github.com/amirfakhrullah/go-bitly/db"
	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUserId(ctx *fiber.Ctx) (uint, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"].(float64)
	if id == 0 {
		return 0, errors.New("invalid user id")
	}
	var userData model.User
	if err := db.DB.Find(&userData, "id = ?", id).Error; err != nil {
		return 0, errors.New(err.Error())
	}
	if userData.ID == 0 {
		return 0, errors.New("user not found")
	}
	return uint(id), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}