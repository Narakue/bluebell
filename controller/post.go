package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"strconv"
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
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ResponseError(c, CodeParam)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ResponseError(c, CodeParam)
		return
	}
	postList, err := logic.GetPostList(page, pageSize)
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
		return
	}
	ResponseSuccess(c, postList)
}

func GetPostDetailByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeParam)
		return
	}
	apiPostDetail, err := logic.GetPostDetailByID(id)
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
		return
	}
	ResponseSuccess(c, apiPostDetail)
}
