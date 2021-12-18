package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/anodoko/service"
	"github.com/labstack/echo/v4"
)


// HandleGetCommentsPageList GET /comics 漫画のシリーズ一覧の取得
func (ctrl Controller) HandleGetCommentsPageList(c echo.Context) error {
	var s service.CommentService
	p, err := s.GetCommentsPageList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}