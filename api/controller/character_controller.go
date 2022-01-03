package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/anodoko/service"
	"github.com/labstack/echo/v4"
)

// HandleGetCharacterUnderVolList GET /characters/under/:cid/:volume
func (ctrl Controller) HandleGetCharacterUnderVolList(c echo.Context) error {
	var s service.CharacterService
	p, err := s.GetCharacterUnderVolList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetCharacterId GET /characters/name/:name キャラ名からIDを取得
func (ctrl Controller) HandleGetCharacterId(c echo.Context) error {
	var s service.CharacterService
	p, err := s.GetCharacterId(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostAppearanceCharacters POST /characters/appearance/ すでに登録されている登場キャラがどのシーンに登場しているか追加
func (ctrl Controller) HandlePostAppearanceCharacters(c echo.Context) error {
	var s service.CharacterService
	p, err := s.PostAppearanceCharacters(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
} 

// HandlePostCharacter POST /characters/ キャラの追加
func (ctrl Controller) HandlePostCharacter(c echo.Context) error {
	var s service.CharacterService
	p, err := s.PostCharacter(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}