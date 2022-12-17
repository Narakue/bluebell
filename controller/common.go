package controller

import (
	"bluebell/logic"
	"bluebell/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func RefreshToken(c *gin.Context) {
	rToken := c.Request.Header.Get("Authorization")
	aToken := c.Query("a_token")
	if len(rToken) == 0 || len(aToken) == 0 {
		ResponseError(c, CodeNotLogin)
		return
	}
	rToken = strings.Fields(rToken)[1]
	// 校验token
	_, err := util.ParseToken(rToken)
	if err != nil {
		ResponseError(c, CodeAuthError)
		return
	}
	_, err = util.ParseToken(aToken)
	if err != nil {
		ResponseError(c, CodeAuthError)
		return
	}
	claims, err := util.ParseToken(aToken)
	if err != nil {
		ResponseError(c, CodeAuthError)
		return
	}
	var token string
	token, err = logic.RefreshToken(claims.UserID)
	if err != nil {
		ResponseError(c, CodeAuthError)
		return
	}
	ResponseSuccess(c, gin.H{"access_token": token})
}
