package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"time"
)

type CommentService struct{}

type CommentsPageList struct{
	Uid  				int
	Name 				string
	Mail 				string
	Image 			string
	Scid 				int
	Page 				int
	Scene 			string
	Comment 		string
	Emotion 		string
	DetailScene string
	CreatedAt 	time.Time
}

// cvidとpageを指定して指定したコメント一覧を取得
func (s CommentService) GetCommentsPageList(db *gorm.DB, c echo.Context) ([]CommentsPageList, error) {
	var cpl []CommentsPageList
	cvid := c.Param("cvid")
	page := c.Param("page")

	if err := db.Debug().Raw("SELECT uid, users.name, users.mail, users.image, scid, page, scene, comment, emotion, detail_scene, scenes.created_at FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid JOIN comics ON comics.id = comic_volumes.cid JOIN users ON scenes.uid = users.id WHERE comic_volumes.id = ? AND scenes.page = ? GROUP BY scid, comment", cvid, page).Scan(&cpl).Error; err != nil {
		return nil, err
	}
	return cpl, nil
}