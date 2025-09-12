package v1

import (
	"net/http"

	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/ctl"
	"github.com/cerami-craft-shop/merchant-backend/repository/db/dao"
	"github.com/cerami-craft-shop/merchant-backend/service"
	"github.com/gin-gonic/gin"
)

// GetItemListHandler godoc
// @Summary 获取商品列表
// @Description 获取所有商品的列表信息
// @Tags 商品管理
// @Accept json
// @Produce json
// @Success 200 {object} ctl.Response{data=types.ItemListResp} "成功获取商品列表"
// @Failure 500 {object} ctl.Response "获取商品列表失败"
// @Router /api/v1/items [get]
func GetItemListHandler(ctx *gin.Context)  {
	itemList, err := service.GetItemService(dao.GetItemDao()).GetItemList()
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, itemList))
}
