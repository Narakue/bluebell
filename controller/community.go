package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCommunityList(c *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, communityList)
}

func GetCommunityDetailByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeParam)
		return
	}
	community, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		ResponseWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, community)
}
