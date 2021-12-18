package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/anodoko/service"
	"github.com/labstack/echo/v4"
)

// HandleGetComicList GET /comics 漫画のシリーズ一覧の取得
func (ctrl Controller) HandleGetComicList(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComicList(ctrl.Db)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostComic POST /comics 漫画シリーズの追加
func (ctrl Controller) HandlePostComic(c echo.Context) error {
	var s service.ComicService
	p, err := s.PostComic(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetComicVolumes GET /comics/volumes/:cid cidの漫画の登録済みの巻一覧を取得
func (ctrl Controller) HandleGetComicVolumes(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComicVolumes(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostComicVolume POST /comics/volume/:cid/:volume 漫画の巻の登録
func (ctrl Controller) HandlePostComicVolume(c echo.Context) error {
	var s service.ComicService
	p, err := s.PostComicVolume(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetComic GET /comics/id 指定されたIDの漫画の取得
func (ctrl Controller) HandleGetComic(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComic(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetComicVolume GET comics/volume/:id 指定された漫画の巻情報の取得
func (ctrl Controller) HandleGetComicVolume(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComicVolume(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetComicFirstVolumeList GET /comics/first/volumes
func (ctrl Controller) HandleGetComicFirstVolumeList(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComicFirstVolumeList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetComicUnderVolumeList GET /comics/under/:cid/:volume
func (ctrl Controller) HandleGetComicUnderVolumeList(c echo.Context) error {
	var s service.ComicService
	p, err := s.GetComicUnderVolumeList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}