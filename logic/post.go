package logic

import (
	"bluebell/dao"
	"bluebell/models"
	"bluebell/util"
	"strconv"
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
	for _, v := range postList {
		v.Score, err = GetScoreByPostID(v.PostID)
		if err != nil {
			return postList, err
		}
	}
	return postList, err
}

func GetPostListFilter(postListParam *models.PostListParam) (postList []*models.Post, err error) {
	postList, err = dao.GetPostListFilter(postListParam)
	return
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

func GetScoreByPostID(id int64) (score int, err error) {
	_, err = GetPostDetailByID(id)
	if err != nil {
		return
	}
	score, err = dao.GetScoreByPostID(id)
	if score < 0 {
		score = 0
	}
	return
}

func VotePost(userID int64, vote *models.VoteParam) (err error) {
	_, err = GetPostDetailByID(vote.PostID)
	if err != nil {
		return
	}
	err = dao.VotePost(strconv.FormatInt(userID, 10), strconv.FormatInt(vote.PostID, 10), float64(*vote.Status))
	return
}
