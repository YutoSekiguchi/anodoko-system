package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CharacterService struct{}

type AppearanceCharacter struct {
	Name       string
	Chid       int
}

// cidとvolumeを指定してそこまでに登場した漫画キャラクター一覧の取得
func (s CharacterService) GetCharacterUnderVolList(db *gorm.DB, c echo.Context) ([]AppearanceCharacter, error) {
	var apch []AppearanceCharacter
	cid := c.Param("cid")
	volume := c.Param("volume")

	if err := db.Debug().Raw("SELECT DISTINCT chid, name FROM `appearance_characters` JOIN characters ON appearance_characters.chid = characters.id JOIN comic_volumes ON appearance_characters.cvid = comic_volumes.id WHERE cid = ? AND volume <= ?", cid, volume).Scan(&apch).Error; err != nil {
		return nil, err
	}
	return apch, nil
}