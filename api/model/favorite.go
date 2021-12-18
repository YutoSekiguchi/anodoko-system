package model

import (
	"time"
)

type Favorite struct {
	ID          int       `gorm:"column:id"`
	UID         int       `gorm:"column:uid"`
	ScID        int       `gorm:"column:scid"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}