package services

import (
	"errors"

	"github.com/amirfakhrullah/go-bitly/db"
	"github.com/amirfakhrullah/go-bitly/model"
)

func GetAllLinks() ([]model.Link, error) {
	var links []model.Link

	if err := db.DB.Find(&links).Error; err != nil {
		return []model.Link{}, err
	}
	return links, nil
}

func GetLink(id uint64) (model.Link, error) {
	var link model.Link
	if err := db.DB.Find(&link, "id = ?", id).Error; err != nil {
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

func UpdateLink(updatedLink model.Link, id uint64) error {
	var link model.Link
	db.DB.Find(&link, "id = ?", id)
	link.RedirectUrl = updatedLink.RedirectUrl
	link.ShortenedId = updatedLink.ShortenedId
	link.Clicked = updatedLink.Clicked
	tx := db.DB.Save(&link)
	return tx.Error
}

func DeleteLink(id uint64) error {
	tx := db.DB.Delete(&model.Link{}, "id = ?", id)
	return tx.Error
}

func FindByShortenedId(id string) (model.Link, error) {
	var l model.Link
	tx := db.DB.Find(&l, "shortened_id = ?", id)
	return l, tx.Error
}
