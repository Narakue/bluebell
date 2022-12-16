package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  GetCodeMsg(CodeSuccess),
		Data: data,
	})
}

func ResponseError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  GetCodeMsg(code),
		Data: nil,
	})
}

func ResponseWithMsg(c *gin.Context, code int, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
