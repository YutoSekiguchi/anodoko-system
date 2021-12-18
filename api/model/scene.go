package model

import (
	"time"
)

type Scene struct {
	ID          int       `gorm:"column:id"`
	UID         int       `gorm:"column:uid"`
	CVID        int       `gorm:"column:cvid"`
	Page        int       `gorm:"column:page"`
	Scene       string    `gorm:"column:scene"`
	Emotion     string    `gorm:"column:emotion"`
	DetailScene string    `gorm:"column:detail_scene"`
	Comment     string    `gorm:"column:comment"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}