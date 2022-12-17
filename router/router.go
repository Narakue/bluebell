package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
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
