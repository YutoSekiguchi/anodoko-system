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