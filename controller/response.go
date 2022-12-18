package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
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

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  GetCodeMsg(code),
		Data: nil,
	})
}

func ResponseWithMsg(c *gin.Context, code ResCode, msg error) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg.Error(),
		Data: nil,
	})
}
