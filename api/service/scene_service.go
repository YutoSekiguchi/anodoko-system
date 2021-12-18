package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"time"
)

type SceneService struct{}

type SceneInfo struct {
	ID          	int       `gorm:"column:id"`
	UID         	int       `gorm:"column:uid"`
	CVID        	int       `gorm:"column:cvid"`
	Page        	int       `gorm:"column:page"`
	Scene       	string    `gorm:"column:scene"`
	Emotion     	string    `gorm:"column:emotion"`
	DetailScene 	string    `gorm:"column:detail_scene"`
	Comment     	string    `gorm:"column:comment"`
	CreatedAt   	time.Time `gorm:"column:created_at"`
	EbookUrl    	string
	ComicVolImage string
}

//GetSceneInfo cid，volume，pageを指定して指定したページに存在するシーンの取得
func (s SceneService) GetSceneInfo(db *gorm.DB, c echo.Context) ([]SceneInfo, error) {
	var sci []SceneInfo
	cid := c.Param("cid")
	volume := c.Param("volume")
	page := c.Param("page")

	if err := db.Debug().Raw("SELECT scenes.id, scenes.uid, scenes.cvid, scenes.page, scenes.scene, scenes.emotion, scenes.detail_scene, scenes.comment, scenes.created_at, comic_volumes.ebook_url, comic_volumes.comic_vol_image FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid JOIN comics ON comics.id = comic_volumes.cid WHERE comic_volumes.cid = ? AND comic_volumes.volume <= ? AND scenes.page = ?", cid, volume, page).Scan(&sci).Error; err != nil {
		return sci, err
	}
	return sci, nil
}


type ScenePages struct {
	Cid    int
	Volume int
	Page   int
	Countpage  int
}
// GetScenePagesASC chid1~5とcid,volume,sceneを指定して検索に応じたシーンの取得（マイナー順）
func (s SceneService) GetScenePagesAsc(db *gorm.DB, c echo.Context) ([]ScenePages, error) {
	var sp []ScenePages
	chid1 := c.Param("chid1")
	chid2 := c.Param("chid2")
	chid3 := c.Param("chid3")
	chid4 := c.Param("chid4")
	chid5 := c.Param("chid5")
	cid := c.Param("cid")
	volume := c.Param("volume")
	scene := c.Param("scene")

	if err := db.Debug().Raw("SELECT cid, volume, page, COUNT(page) as countpage FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid WHERE chid IN (?, ?, ?, ?, ?) AND comic_volumes.cid = ? AND comic_volumes.volume <= ? AND scenes.scene = ? GROUP BY cid, volume, page ORDER BY COUNT(page) ASC", chid1, chid2, chid3, chid4, chid5, cid, volume, scene).Scan(&sp).Error; err != nil {
		return nil, err
	}
	return sp, nil
}

// GetScenePagesDESC chid1~5とcid,volume,sceneを指定して検索に応じたシーンの取得（人気順）
func (s SceneService) GetScenePagesDesc(db *gorm.DB, c echo.Context) ([]ScenePages, error) {
	var sp []ScenePages
	chid1 := c.Param("chid1")
	chid2 := c.Param("chid2")
	chid3 := c.Param("chid3")
	chid4 := c.Param("chid4")
	chid5 := c.Param("chid5")
	cid := c.Param("cid")
	volume := c.Param("volume")
	scene := c.Param("scene")

	if err := db.Debug().Raw("SELECT cid, volume, page, COUNT(page) as countpage FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid WHERE chid IN (?, ?, ?, ?, ?) AND comic_volumes.cid = ? AND comic_volumes.volume <= ? AND scenes.scene = ? GROUP BY cid, volume, page ORDER BY COUNT(page) DESC", chid1, chid2, chid3, chid4, chid5, cid, volume, scene).Scan(&sp).Error; err != nil {
		return nil, err
	}
	return sp, nil
}

// GetScenesVolList cid, volumeを指定して指定した漫画の巻のシーンを取得
func (s SceneService) GetScenesVolList(db *gorm.DB, c echo.Context) ([]ScenePages, error) {
	var sp []ScenePages
	cid := c.Param("cid")
	volume := c.Param("volume")

	if err := db.Debug().Raw("SELECT cid, volume, page, COUNT(page) as countpage FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid WHERE comic_volumes.cid = ? AND comic_volumes.volume = ? GROUP BY cid, volume, page", cid, volume).Scan(&sp).Error; err != nil {
		return nil, err
	}
	return sp, nil
}


type ScenePlusVolume struct {
	ID          	int       `gorm:"column:id"`
	UID         	int       `gorm:"column:uid"`
	CVID        	int       `gorm:"column:cvid"`
	Page        	int       `gorm:"column:page"`
	Scene       	string    `gorm:"column:scene"`
	Emotion     	string    `gorm:"column:emotion"`
	DetailScene 	string    `gorm:"column:detail_scene"`
	Comment     	string    `gorm:"column:comment"`
	CreatedAt   	time.Time `gorm:"column:created_at"`
	Volume      	int
	ComicVolImage string
	EbookUrl      string
}
// GetScenesUnderVolList chid1~5とcid, volume, sceneを指定してそれに合うシーンを全て取得
func (s SceneService) GetScenesUnderVolList(db *gorm.DB, c echo.Context) ([]ScenePlusVolume, error) {
	var spv []ScenePlusVolume
	chid1 := c.Param("chid1")
	chid2 := c.Param("chid2")
	chid3 := c.Param("chid3")
	chid4 := c.Param("chid4")
	chid5 := c.Param("chid5")
	cid := c.Param("cid")
	volume := c.Param("volume")
	scene := c.Param("scene")

	if err := db.Debug().Raw("SELECT *, comic_volumes.volume, comic_volumes.ebook_url FROM `scenes` JOIN appearance_characters ON appearance_characters.scid = scenes.id JOIN characters ON characters.id = appearance_characters.chid JOIN comic_volumes ON comic_volumes.id = appearance_characters.cvid WHERE chid IN (?, ?, ?, ?, ?) AND comic_volumes.cid = ? AND comic_volumes.volume <= ? AND scenes.scene = ?", chid1, chid2, chid3, chid4, chid5, cid, volume, scene).Scan(&spv).Error; err != nil {
		return nil, err
	}
	return spv, nil
}

type SceneUserList struct {
	ID          	int       `gorm:"column:id"`
	UID         	int       `gorm:"column:uid"`
	CVID        	int       `gorm:"column:cvid"`
	Page        	int       `gorm:"column:page"`
	Scene       	string    `gorm:"column:scene"`
	Emotion     	string    `gorm:"column:emotion"`
	DetailScene 	string    `gorm:"column:detail_scene"`
	Comment     	string    `gorm:"column:comment"`
	CreatedAt   	time.Time `gorm:"column:created_at"`
	Name          string
	Volume      	int
	ComicVolImage string
	EbookUrl      string
}

// uidを指定してユーザが追加したシーンの取得
func (s SceneService) GetUserSetSceneList(db *gorm.DB, c echo.Context) ([]SceneUserList, error) {
	var sul []SceneUserList
	uid := c.Param("uid")

	if err := db.Debug().Raw("SELECT *, comics.name, comic_volumes.volume, comic_volumes.ebook_url FROM `scenes` JOIN comic_volumes ON comic_volumes.id = scenes.cvid JOIN comics ON comics.id = comic_volumes.cid WHERE scenes.uid = ? ", uid).Scan(&sul).Error; err != nil {
		return nil, err
	}
	return sul, nil
}