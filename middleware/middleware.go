package middleware

import (
	"bluebell/controller"
	"bluebell/dao/redis"
	"bluebell/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var token string

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			controller.ResponseError(c, controller.CodeNotLogin)
			c.Abort()
		}
		auth = strings.Fields(auth)[1]
		token = auth
		// 校验token
		claim, err := util.ParseToken(auth)
		if err != nil {
			controller.ResponseError(c, controller.CodeAuthError)
			c.Abort()
		}
		c.Set(util.CtxUserID, claim.UserID)
		c.Next()
	}
}

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := controller.GetUserID(c)
		if !exists {
			controller.ResponseError(c, controller.CodeNotLogin)
			c.Abort()
		}
		rdb := redis.GetRdb()
		rdbToken := rdb.Get(strconv.FormatInt(value, 10) + util.AToken)
		if token != rdbToken.Val() {
			controller.ResponseError(c, controller.CodeReLogin)
			c.Abort()
		}
		c.Next()
	}
}
