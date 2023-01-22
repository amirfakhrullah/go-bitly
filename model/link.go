package model

func GetAllLinks() ([]Link, error) {
	var links []Link

	tx := db.Find(&links)
	if tx.Error != nil {
		return []Link{}, tx.Error
	}
	return links, nil
}

func GetLink(id uint64) (Link, error) {
	var link Link
	tx := db.Where("id = ?", id).First(&link)
	return link, tx.Error
}

func CreateLink(l Link) error {
	tx := db.Create(&l)
	return tx.Error
}

func UpdateLink(updatedLink Link, id uint64) error {
	var link Link
	db.Where("id = ?", id).First(&link)
	link.RedirectUrl = updatedLink.RedirectUrl
	link.ShortenedId = updatedLink.ShortenedId
	link.Clicked = updatedLink.Clicked
	tx := db.Save(&link)
	return tx.Error
}

func DeleteLink(id uint64) error {
	tx := db.Unscoped().Delete(&Link{}, id)
	return tx.Error
}

func FindByShortenedId(id string) (Link, error) {
	var l Link
	tx := db.Where("shortened_id = ?", id).First(&l)
	return l, tx.Error
}