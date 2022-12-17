package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("param error", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeParam)
			return
		}
		ResponseWithMsg(c, CodeParam, errs)
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
	aToken, rToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeLogin)
		return
	}
	result := map[string]any{
		"access_token":  aToken,
		"refresh_token": rToken,
	}
	ResponseSuccess(c, result)
}

func GetUserID(c *gin.Context) (int64, bool) {
	value, exists := c.Get(util.CtxUserID)
	if !exists {
		return 0, false
	} else {
		return value.(int64), true
	}
}
