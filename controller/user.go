package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("param error", zap.Error(err))
		ResponseError(c, CodeParam)
		return
	}

	if err := logic.SignUp(p); err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseWithMsg(c, CodeSignUp, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func Login(c *gin.Context) {
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeParam)
		return
	}
	if err := logic.Login(p); err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeLogin)
		return
	}
	ResponseSuccess(c, nil)
}
