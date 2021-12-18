package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ComicService struct{}

// 漫画シリーズの一覧を取得
func (s ComicService) GetComicList(db *gorm.DB) ([]Comic, error) {
	var cm []Comic

	if err := db.Find(&cm).Error; err != nil {
		return nil, err
	}
	return cm, nil
}

// PostComic 漫画のシリーズの追加
// request bodyでjsonを渡す
func (s ComicService) PostComic(db *gorm.DB, c echo.Context) (Comic, error) {
	var cm Comic
	c.Bind(&cm)

	if err := db.Create(&cm).Error; err != nil {
		return cm, err
	}
	return cm, nil
}

// idを指定してその漫画を抽出
func (s ComicService) GetComic(db *gorm.DB, c echo.Context) ([]Comic, error) {
	var cm []Comic
	id := c.Param("id")

	if err := db.Debug().Raw("SELECT * FROM comics WHERE id = ?", id).Scan(&cm).Error; err != nil {
		return nil, err
	}
	return cm, nil
}

// GetComicVolume cvidを指定して漫画の取得
func (s ComicService) GetComicVolume(db *gorm.DB, c echo.Context) ([]Volume, error) {
	var v []Volume
	cid := c.Param("cid")
	volume := c.Param("volume")

	if err := db.Debug().Raw("SELECT * FROM comic_volumes WHERE cid = ? AND volume = ?", cid, volume).Scan(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

// ある漫画シリーズの巻一覧の取得
func (s ComicService) GetComicVolumes(db *gorm.DB, c echo.Context) ([]Volume, error) {
	var v []Volume
	cid := c.Param("cid")

	if err := db.Debug().Raw("SELECT * FROM comic_volumes WHERE cid = ? ORDER BY volume ASC", cid).Scan(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

// PostComicVolume 漫画の巻の追加
func (s ComicService) PostComicVolume(db *gorm.DB, c echo.Context) (Volume, error) {
	var v Volume
	c.Bind(&v)

	if err := db.Table("comic_volumes").Create(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

// GetComicFirstVolumeList 追加されてる漫画の最初の巻の追加
func (s ComicService) GetComicFirstVolumeList(db *gorm.DB, c echo.Context) ([]Volume, error) {
	var v []Volume
	
	if err := db.Debug().Raw("SELECT * FROM (SELECT *, row_number()over(PARTITION BY cid ORDER BY volume) row_num FROM comic_volumes) AS volume_image WHERE row_num = 1").Scan(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

// GetComicUnderVolumeList cidとvolumeを指定してその漫画のその巻以下の漫画の取得
func (s ComicService) GetComicUnderVolumeList(db *gorm.DB, c echo.Context) ([]Volume, error) {
	var v []Volume
	cid := c.Param("cid")
	volume := c.Param("volume")

	if err := db.Debug().Raw("SELECT * FROM comic_volumes WHERE cid = ? AND volume <= ? ", cid, volume).Scan(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}