package dao

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"errors"
	"strconv"
	"time"
)

const (
	onWeekForInt = 7 * 24 * 3600
	score        = 432
)

func CreatePost(post *models.Post) (bool, error) {
	db := mysql.GetDB()
	rdb := redis.GetRdb()
	rdb.ZAdd(redis.GetKey(redis.KeyPostVoteTime), redis.MakeZ(float64(time.Now().Unix()), post.PostID))
	result := db.Create(post)
	return result.RowsAffected > 0, result.Error
}

func GetPostList(page, pageSize int) ([]*models.Post, error) {
	db := mysql.GetDB()
	postList := make([]*models.Post, 0)
	result := db.Limit(pageSize).Offset(page).Find(&postList)
	return postList, result.Error
}

func GetPostDetailByID(id int64) (*models.Post, error) {
	db := mysql.GetDB()
	post := &models.Post{PostID: id}
	result := db.Where(post).First(post)
	return post, result.Error
}

func GetScoreByPostID(id int64) (score int, err error) {
	rdb := redis.GetRdb()
	score = int(rdb.ZScore(redis.GetKey(redis.KeyPostVoteScore), strconv.FormatInt(id, 10)).Val())
	return
}

func VotePost(userID, postID string, status float64) (err error) {
	rdb := redis.GetRdb()
	voteTime := rdb.ZScore(redis.GetKey(redis.KeyPostVoteTime), postID).Val()
	if float64(time.Now().Unix())-voteTime > onWeekForInt {
		// 帖子时间超过一周
		err = errors.New("you can't do that")
		return
	}
	voteUserStatus := rdb.ZScore(redis.GetKey(redis.KeyPostVoteUser)+userID, postID).Val()
	curScore := rdb.ZScore(redis.GetKey(redis.KeyPostVoteScore), postID).Val()
	curScore = (status-voteUserStatus)*score + curScore
	rdb.ZAdd(redis.GetKey(redis.KeyPostVoteScore), redis.MakeZ(curScore, postID))
	rdb.ZAdd(redis.GetKey(redis.KeyPostVoteUser)+userID, redis.MakeZ(status, postID))
	return
}
