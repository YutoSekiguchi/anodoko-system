package model

import "time"

type Volume struct {
	ID            int       `gorm:"column:id"`
	Volume        int       `gorm:"column:volume"`
	CID           int       `gorm:"column:cid"`
	AllPage       int       `gorm:"column:all_page"`
	EbookURL      string    `gorm:"column:ebook_url"`
	ComicVolImage string    `gorm:"column:comic_vol_image"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}