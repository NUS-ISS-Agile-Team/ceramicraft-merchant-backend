package main

import (
	"fmt"

	"github.com/cerami-craft-shop/merchant-backend/config"
	"github.com/cerami-craft-shop/merchant-backend/repository/db"
	"github.com/cerami-craft-shop/merchant-backend/router"
)

func main() {
	initAll()
	r := router.NewRouter()
	_ = r.Run(fmt.Sprintf("%s:%s", config.Config.System.Host, config.Config.System.Port))
}

func initAll() {
	config.InitConfig()
	db.Init()
}
