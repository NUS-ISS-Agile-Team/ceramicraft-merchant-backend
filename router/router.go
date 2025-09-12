package router

import (
	api "github.com/cerami-craft-shop/merchant-backend/api/v1"
	"github.com/gin-gonic/gin"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/cerami-craft-shop/merchant-backend/docs"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// swagger router
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})

		v1.GET("item/list", api.GetItemListHandler)
	}
	return r
}
