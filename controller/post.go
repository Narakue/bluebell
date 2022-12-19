package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CreatePost(c *gin.Context) {
	postParam := new(models.PostParam)
	err := c.ShouldBindJSON(postParam)
	if err != nil {
		fmt.Println(err)
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
	page, pageSize, err := getPageInfo(c)
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

func GetPostListFilter(c *gin.Context) {
	p := &models.PostListParam{
		Page:     0,
		PageSize: 10,
		Order:    models.OrderTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		ResponseError(c, CodeParam)
	}
	list, err := logic.GetPostListFilter(p)
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
		return
	}
	ResponseSuccess(c, list)
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

func VotePost(c *gin.Context) {
	voteParam := new(models.VoteParam)
	err := c.ShouldBindJSON(voteParam)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeParam)
		return
	}
	userID, ok := GetUserID(c)
	if !ok {
		ResponseError(c, CodeReLogin)
		return
	}
	err = logic.VotePost(userID, voteParam)
	if err != nil {
		ResponseWithMsg(c, CodeError, err)
		return
	}
	ResponseSuccess(c, nil)
}
