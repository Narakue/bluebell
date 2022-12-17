package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
)

func GetCommunityList(c *gin.Context) {
	communityList := logic.GetCommunityList()
	ResponseSuccess(c, communityList)
}
