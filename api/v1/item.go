package v1

import (
	"net/http"

	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/ctl"
	"github.com/cerami-craft-shop/merchant-backend/service"
	"github.com/gin-gonic/gin"
)

func GetItemListHandler(ctx *gin.Context)  {
	itemList, err := service.GetItemService().GetItemList()
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, itemList))
}
