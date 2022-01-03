package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
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

// nameを指定してそのキャラクターのidを取得
func (s CharacterService) GetCharacterId(db *gorm.DB, c echo.Context) ([]Character, error) {
	var ch []Character
	name := c.Param("name")
	
	if err := db.Debug().Raw("SELECT * FROM characters WHERE name = ?", name).Scan(&ch).Error; err != nil {
		return nil, err
	}
	return ch, nil
}

type PostAppearanceCharacter struct {
	Id         int
	Scid       int
	Cvid       int
	Chid       int
	CreatedAt  time.Time
}

// POST
// PostAppearanceCharacters
func (s CharacterService) PostAppearanceCharacters(db *gorm.DB, c echo.Context) (PostAppearanceCharacter, error) {
	var pac PostAppearanceCharacter
	c.Bind(&pac)

	if err := db.Table("appearance_characters").Create(&pac).Error; err != nil {
		return pac, err
	}
	return pac, nil
}

// PostCharacter
func (s CharacterService) PostCharacter(db *gorm.DB, c echo.Context) (Character, error) {
	var ch Character
	c.Bind(&ch)

	if err := db.Table("characters").Create(&ch).Error; err != nil {
		return ch, err
	}
	return ch, nil
}