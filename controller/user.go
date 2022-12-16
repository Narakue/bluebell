package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	p := new(models.SignUpParam)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("param error", zap.Error(err))
		return
	}
	err = logic.SignUp(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	c.JSON(200, gin.H{"msg": "success"})
}
