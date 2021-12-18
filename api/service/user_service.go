package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct{}

func (s UserService) GetUserList(db *gorm.DB) ([]User, error) {
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// PostUser ユーザ追加
// request bodyでjsonを渡す
func (s UserService) PostUser(db *gorm.DB, c echo.Context) (User, error) {
	var user User
	c.Bind(&user)

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

type UserName struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

// GetUserByID uidが一致するユーザを取得
func (s UserService) GetUserByID(db *gorm.DB, c echo.Context) (UserName, error) {
	var u UserName
	uid := c.Param("uid")

	if err := db.Table("users").Where("id = ?", uid).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// CheckUserByEmail そのメールアドレスのユーザが存在するかどうか
// param: json { "Mail": "target@dummy.com" }
func (s UserService) GetUserByEmail(db *gorm.DB, c echo.Context) (User, error) {
	var u User
	mail := c.QueryParam("Mail")
	fmt.Println(mail)
	if err := db.Where("mail = ?", mail).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}