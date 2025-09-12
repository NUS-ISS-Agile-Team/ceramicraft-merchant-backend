package ctl

import (
	"github.com/cerami-craft-shop/merchant-backend/pkg/consts"
	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := consts.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    consts.GetMsg(status),
	}

	return r
}

// RespError 错误返回
func RespError(ctx *gin.Context, err error, code ...int) *Response {
	status := consts.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Msg:    consts.GetMsg(status),
		Data:   nil,
		Error:  err.Error(),
	}

	return r
}
