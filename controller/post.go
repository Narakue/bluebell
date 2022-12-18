package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"time"
)

func CreatePost(c *gin.Context) {
	postParam := new(models.PostParam)
	err := c.ShouldBindJSON(postParam)
	if err != nil {
		ResponseError(c, CodeParam)
		return
	}
	userID, ok := GetUserID(c)
	if !ok {
		ResponseError(c, CodeServerBusy)
		return
	}
	post := &models.Post{
		Title:       postParam.Title,
		Content:     postParam.Content,
		AuthorID:    userID,
		CommunityID: postParam.CommunityID,
		Status:      models.StatusExist,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	err = logic.CreatePost(post)
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostList(c *gin.Context) {
	postList, err := logic.GetPostList()
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
	}
	ResponseSuccess(c, postList)
}
