package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FavoriteService struct{}

// uidとscidを指定してそのシーンがお気に入りに入っていれば抽出
func (s FavoriteService) GetWhetherFavorite(db *gorm.DB, c echo.Context) ([]Favorite, error) {
	var f []Favorite
	uid := c.Param("uid")
	scid := c.Param("scid")

	if err := db.Where("uid = ? AND scid = ?", uid, scid).Find(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

type FavoriteSceneInfo struct {
	Id        int
	Name 				string
	Volume			int
	DetailScene string
	Page 				int
}

// uidを指定して指定したユーザーのお気に入りの一覧を取得
func (s FavoriteService) GetFavoritesList(db *gorm.DB, c echo.Context) ([]FavoriteSceneInfo, error) {
	var fsi []FavoriteSceneInfo
	uid := c.Param("uid")

	if err := db.Debug().Raw("SELECT DISTINCT scenes.id, comics.name, comic_volumes.volume, scenes.detail_scene, scenes.page FROM favorites JOIN scenes ON scenes.id = favorites.scid JOIN comic_volumes ON comic_volumes.id = scenes.cvid JOIN comics ON comics.id = comic_volumes.cid WHERE favorites.uid = ?", uid).Scan(&fsi).Error; err != nil {
		return nil, err
	}
	return fsi, nil
}

// POST
// お気に入りの追加
func (s FavoriteService) PostFavorites(db *gorm.DB, c echo.Context) (Favorite, error) {
	var f Favorite
	c.Bind(&f)

	if err := db.Table("favorites").Create(&f).Error; err != nil {
		return f, err
	}
	return f, nil
}

// DELETE
// お気に入りの削除
func (s FavoriteService) DeleteFavorites(db *gorm.DB, c echo.Context) ([]Favorite, error) {
	var f []Favorite
	uid := c.Param("uid")
	scid := c.Param("scid")

	if err := db.Where("uid = ? AND scid = ?", uid, scid).Delete(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

