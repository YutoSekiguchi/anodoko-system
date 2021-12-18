package model

import (
	"time"
)

type Comic struct {
	ID         int       `gorm:"primary_key;not null;autoIncrement:true"`
	Name       string    `gorm:"type:text;not null; column:name"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp;column:created_at"`
}