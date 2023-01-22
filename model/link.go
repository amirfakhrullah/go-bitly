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

func UpdateLink(l Link) error {
	tx := db.Save(&l)
	return tx.Error
}

func DeleteLink(id uint64) error {
	tx := db.Unscoped().Delete(&Link{}, id)
	return tx.Error
}

func FindByUrl(url string) (Link, error) {
	var l Link
	tx := db.Where("shortenedUrl = ?", url).First(&l)
	return l, tx.Error
}