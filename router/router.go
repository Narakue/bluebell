package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	UserGroup(r)
	return r
}

func UserGroup(r *gin.Engine) {
	user := r.Group("user")
	{
		user.POST("/sign", controller.SignUp)
		user.POST("/login", controller.Login)
	}
}
