package model

type Link struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	RedirectUrl string `json:"redirect_url" gorm:"unique;not null"`
	ShortenedId string `json:"shortened_id" gorm:"unique;not null"`
	Clicked     uint64 `json:"clicked"`
}
