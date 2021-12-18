package model

import (
	"time"
)

// User DBのユーザテーブルの構成
type User struct {
	ID          int       `gorm:"primary_key;not null;autoIncrement:true"`
	Name        string    `gorm:"type:text;not null"`
	Mail        string    `gorm:"type:text;not null"`
	Image       string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `sql:"DEFAULT:current_timestamp;column:created_at"`
}