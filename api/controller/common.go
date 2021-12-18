package controller

import "gorm.io/gorm"

// Controller Struct 以下の関数を構造体のメソッドにするためのもの
type Controller struct {
	Db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		Db: db,
	}
}