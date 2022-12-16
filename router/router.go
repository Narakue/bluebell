package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/sign", controller.SignUp)
	return r
}
