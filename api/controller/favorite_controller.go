package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/anodoko/service"
	"github.com/labstack/echo/v4"
)


// HandleGetWhetherFavorite GET /favorites/whether/:uid/:scid お気に入りシーンかどうか取得
func (ctrl Controller) HandleGetWhetherFavorite(c echo.Context) error {
	var s service.FavoriteService
	p, err := s.GetWhetherFavorite(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetFavoritesList Get /favorites/list/:uid ユーザのお気に入りシーンの一覧
func (ctrl Controller) HandleGetFavoritesList(c echo.Context) error {
	var s service.FavoriteService
	p, err := s.GetFavoritesList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error() )
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostFavorites POST /favorites お気に入りシーンの追加
func (ctrl Controller) HandlePostFavorites(c echo.Context) error {
	var s service.FavoriteService
	p, err := s.PostFavorites(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleDeleteFavorites DELETE /favorites/:uid/:scid お気に入りシーンの削除
func (ctrl Controller) HandleDeleteFavorites(c echo.Context) error {
	var s service.FavoriteService
	p, err := s.DeleteFavorites(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}