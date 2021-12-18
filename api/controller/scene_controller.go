package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/anodoko/service"
	"github.com/labstack/echo/v4"
)

// HandleGetSceneInfo GET /scenes/info/:cid/:volume/:page 指定したページに存在するシーンの取得
func (ctrl Controller) HandleGetSceneInfo(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetSceneInfo(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetScenePagesAsc GET  /scenes/pages/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene/asc 検索結果を昇順で表示
func (ctrl Controller) HandleGetScenePagesAsc(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetScenePagesAsc(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetScenePagesDesc GET  /scenes/pages/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene/asc 検索結果を降順で表示
func (ctrl Controller) HandleGetScenePagesDesc(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetScenePagesDesc(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetScenesVolList GET /scenes/volume/pages/:cid/:volume 指定した巻の登録ページを取得
func (ctrl Controller) HandleGetScenesVolList(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetScenesVolList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetScenesUnderVolList GET /scenes//volume/pages/under/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene 検索結果に合うシーン一覧
func (ctrl Controller) HandleGetScenesUnderVolList(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetScenesUnderVolList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetUserSetSceneList GET /scenes/list/user/:uid ユーザが登録したシーン一覧
func (ctrl Controller) HandleGetUserSetSceneList(c echo.Context) error {
	var s service.SceneService
	p, err := s.GetUserSetSceneList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}