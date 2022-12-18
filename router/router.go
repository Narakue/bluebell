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
	UserGroup(r)
	r.GET("/refreshToken", controller.RefreshToken)
	r.Use(middleware.AuthMiddleware())
	//r.Use(middleware.LoginMiddleware())
	CommunityGroup(r)
	PostGroup(r)
	TestGroup(r)
	return r
}

func UserGroup(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/sign", controller.SignUp)
		user.POST("/login", controller.Login)
	}
}

func CommunityGroup(r *gin.Engine) {
	community := r.Group("/community")
	{
		community.GET("/getCommunityList", controller.GetCommunityList)
		community.GET("/communityDetail/:id", controller.GetCommunityDetailByID)
	}
}

func PostGroup(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.POST("/createPost", controller.CreatePost)
		post.GET("/getPostList", controller.GetPostList)
		post.GET("/postDetail/:id", controller.GetPostDetailByID)
	}
}

func TestGroup(r *gin.Engine) {
	test := r.Group("/test")
	{
		test.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "success"})
		})
	}
}
