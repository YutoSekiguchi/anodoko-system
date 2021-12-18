package main

import (
	"github.com/YutoSekiguchi/anodoko/router"
	"github.com/YutoSekiguchi/anodoko/util"
)

func main() {
	db := util.InitDb() 
	router.InitRouter(db)
}