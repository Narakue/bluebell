package middleware

import (
	"bluebell/controller"
	"bluebell/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			controller.ResponseError(c, controller.COdeAuthError)
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		claim, err := util.ParseToken(auth)
		if err != nil {
			c.Abort()
			controller.ResponseError(c, controller.COdeAuthError)
		}
		c.Set(util.UserID, claim.UserID)
		c.Next()
	}
}
