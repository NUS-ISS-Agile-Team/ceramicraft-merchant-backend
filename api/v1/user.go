package v1

import (
	"net/http"

	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/ctl"
	"github.com/cerami-craft-shop/merchant-backend/pkg/utils/logger"
	"github.com/cerami-craft-shop/merchant-backend/service"
	"github.com/cerami-craft-shop/merchant-backend/types"
	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary 商家端登录
// @Description 商家端登录（只有一个合法帐号）
// @Tags 权限管理
// @Accept json
// @Produce json
// @Param data body types.UserLoginReq true "登录请求体"
// @Success 200 {object} ctl.Response{data=types.UserLoginResp} "登录接口请求成功"
// @Failure 500 {object} ctl.Response "登录接口请求错误"
// @Router /api/v1/user/login [post]
func LoginHandler(ctx *gin.Context) {
	var req types.UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		logger.GetLogger().Infof("login req err, err: %v", err)
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err))
		return
	}

	resp, err := service.GetUserService().Login(req.UserName, req.Password)
	if err != nil {
		logger.GetLogger().Infof("login req err, err: %v", err)
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
