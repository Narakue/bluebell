package logic

import (
	"bluebell/dao"
	"bluebell/models"
	"errors"
)

func CreatePost(post *models.Post) error {
	community, err := GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		return err
	}
	if community != nil {
		return errors.New("community id not exists")
	}
	result, err := dao.CreatePost(post)
	if err != nil || !result {
		return err
	}
	return nil
}

func GetPostList() ([]*models.Post, error) {
	postList, err := dao.GetPostList()
	return postList, err
}
