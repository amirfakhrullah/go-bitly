package services

import (
	"errors"

	"github.com/amirfakhrullah/go-bitly/db"
	"github.com/amirfakhrullah/go-bitly/model"
)

func GetAllLinks(userId uint) ([]model.Link, error) {
	var links []model.Link

	if err := db.DB.Find(&links, "user_id = ?", userId).Error; err != nil {
		return []model.Link{}, err
	}
	return links, nil
}

func GetLink(id uint64, userId uint) (model.Link, error) {
	var link model.Link
	if err := db.DB.Find(&link, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return model.Link{}, err
	}
	if link.ID == 0 {
		return model.Link{}, errors.New("link does not exist")
	}
	return link, nil
}

func CreateLink(l model.Link) error {
	tx := db.DB.Create(&l)
	return tx.Error
}

func UpdateLink(userId uint, updatedLink model.Link, id uint64) error {
	var link model.Link
	db.DB.Find(&link, "id = ? AND user_id = ?", id, userId)
	link.RedirectUrl = updatedLink.RedirectUrl
	link.ShortenedId = updatedLink.ShortenedId
	link.Clicked = updatedLink.Clicked
	tx := db.DB.Save(&link)
	return tx.Error
}

func DeleteLink(userId uint, id uint64) error {
	tx := db.DB.Delete(&model.Link{}, "id = ? AND user_id = ?", id, userId)
	return tx.Error
}

// PUBLIC --------------------
func OpenShortenedId(id string) (model.Link, error) {
	var l model.Link
	if err := db.DB.Find(&l, "shortened_id = ?", id).Error; err != nil {
		return model.Link{}, err
	}
	l.Clicked += 1
	tx := db.DB.Save(&l)
	return l, tx.Error
}
