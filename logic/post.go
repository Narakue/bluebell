package logic

import (
	"bluebell/dao"
	"bluebell/models"
	"bluebell/util"
)

func CreatePost(post *models.Post) error {
	_, err := GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		return err
	}
	post.PostID = util.GenID()
	result, err := dao.CreatePost(post)
	if err != nil || !result {
		return err
	}
	return nil
}

func GetPostList(page, pageSize int) ([]*models.Post, error) {
	postList, err := dao.GetPostList(page, pageSize)
	return postList, err
}

func GetPostDetailByID(id int64) (apiPostDetail *models.ApiPostDetail, err error) {
	post, err := dao.GetPostDetailByID(id)
	if err != nil {
		return
	}
	user, err := GetUserByID(post.AuthorID)
	if err != nil {
		return
	}
	community, err := GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		return
	}
	apiPostDetail = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return
}
