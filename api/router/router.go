package router

import (
	"net/http"


	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/YutoSekiguchi/anodoko/controller"
	// "github.com/YutoSekiguchi/anodoko/util"
)

func InitRouter(db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:7120", "http://vps7.nkmr.io", "https://vps7.nkmr.io", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	ctrl := controller.NewController(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// User
	user := e.Group("/users")
	{
		user.GET("", ctrl.HandleGetUserList)
		user.POST("", ctrl.HandlePostUser)
		user.GET("/:uid", ctrl.HandleGetUserByID)
		user.GET("/me", ctrl.HandleGetUserByEmail)
	}

	// Comic
	comic := e.Group("/comics")
	{
		comic.GET("", ctrl.HandleGetComicList)
		comic.POST("", ctrl.HandlePostComic)
		comic.GET("/:id", ctrl.HandleGetComic)
		comic.GET("/volume/:cid/:volume", ctrl.HandleGetComicVolume)
		comic.GET("/volumes/:cid", ctrl.HandleGetComicVolumes)
		comic.POST("/volumes", ctrl.HandlePostComicVolume)
		comic.GET("/first", ctrl.HandleGetComicFirstVolumeList)
		comic.GET("/under/:cid/:volume", ctrl.HandleGetComicUnderVolumeList)
	}

	// Character
	character := e.Group("/characters")
	{
		character.GET("/under/:cid/:volume", ctrl.HandleGetCharacterUnderVolList)
		character.GET("/name/:name", ctrl.HandleGetCharacterId)
		character.POST("/appearance", ctrl.HandlePostAppearanceCharacters)
		character.POST("", ctrl.HandlePostCharacter)
	}

	// Scene
	scene := e.Group("/scenes")
	{
		scene.GET("/info/:cid/:volume/:page", ctrl.HandleGetSceneInfo)
		scene.GET("/pages/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene/asc", ctrl.HandleGetScenePagesAsc)
		scene.GET("/pages/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene/desc", ctrl.HandleGetScenePagesDesc)
		scene.GET("/volume/pages/:cid/:volume", ctrl.HandleGetScenesVolList)
		scene.GET("/volume/pages/under/:chid1/:chid2/:chid3/:chid4/:chid5/:cid/:volume/:scene", ctrl.HandleGetScenesUnderVolList)
		scene.GET("/list/user/:uid", ctrl.HandleGetUserSetSceneList)
		scene.GET("/scene/:scid", ctrl.HandleGetScene)
		scene.GET("/user/info/:uid/:cvid/:page/:scene", ctrl.HandleGetUserScenesInfo)
		scene.POST("", ctrl.HandlePostScenes)
	}

	// Comment
	comment := e.Group("/comments")
	{
		comment.GET("/list/:cvid/:page", ctrl.HandleGetCommentsPageList)
	}

	// Favorite
	favorite := e.Group("/favorites")
	{
		favorite.GET("/whether/:uid/:scid", ctrl.HandleGetWhetherFavorite)
		favorite.GET("/list/:uid", ctrl.HandleGetFavoritesList)
		favorite.POST("", ctrl.HandlePostFavorites)
		favorite.DELETE("/:uid/:scid", ctrl.HandleDeleteFavorites)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}