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
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(redis.GetKey(redis.KeyPostVoteTime), redis.MakeZ(float64(time.Now().Unix()), post.PostID))
	pipeline.ZAdd(redis.GetKey(redis.KeyPostVoteScore), redis.MakeZ(float64(time.Now().Unix()), post.PostID))
	_, err := pipeline.Exec()
	if err != nil {
		return false, err
	}
	result := db.Create(post)
	return result.RowsAffected > 0, result.Error
}

func GetPostList(page, pageSize int) ([]*models.Post, error) {
	db := mysql.GetDB()
	postList := make([]*models.Post, 0)
	result := db.Limit(pageSize).Offset(page).Order("id desc").Find(&postList)
	return postList, result.Error
}

func GetPostListFilter(postListParam *models.PostListParam) (postList []*models.Post, err error) {
	rdb := redis.GetRdb()
	postList = make([]*models.Post, 0)
	key := ""
	if postListParam.Order == models.OrderTime {
		key = redis.KeyPostVoteTime
	} else {
		key = redis.KeyPostVoteScore
	}
	result, err := rdb.ZRevRange(redis.GetKey(key), int64(postListParam.Page), int64(postListParam.Page+postListParam.PageSize)).Result()
	if err != nil {
		return
	}
	for _, v := range result {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return postList, err
		}
		post, err := GetPostDetailByID(id)
		if err != nil {
			return postList, err
		}
		postList = append(postList, post)
	}
	return
}

func GetPostDetailByID(id int64) (post *models.Post, err error) {
	db := mysql.GetDB()
	post = &models.Post{PostID: id}
	result := db.Where(post).First(post)
	err = result.Error
	post.Score, err = GetScoreByPostID(id)
	return
}

func GetScoreByPostID(id int64) (score int, err error) {
	rdb := redis.GetRdb()
	score = int(rdb.ZScore(redis.GetKey(redis.KeyPostVoteScore), strconv.FormatInt(id, 10)).Val())
	return
}

func VotePost(userID, postID string, status float64) (err error) {
	rdb := redis.GetRdb()
	db := mysql.GetDB()
	pipeline := rdb.TxPipeline()
	voteTime := rdb.ZScore(redis.GetKey(redis.KeyPostVoteTime), postID).Val()
	if float64(time.Now().Unix())-voteTime > onWeekForInt {
		// 帖子时间超过一周
		pipeline.ZRem(redis.GetKey(redis.KeyPostVoteScore), postID)
		pipeline.ZRem(redis.GetKey(redis.KeyPostVoteUser)+postID, userID)
		pipeline.ZRem(redis.GetKey(redis.KeyPostVoteTime), postID)
		pID, err := strconv.ParseInt(postID, 10, 64)
		if err != nil {
			return err
		}
		post, err := GetPostDetailByID(pID)
		if err != nil {
			return err
		}
		db.Save(post)
		err = errors.New("you can't do that")
		return err
	}
	voteUserStatus := pipeline.ZScore(redis.GetKey(redis.KeyPostVoteUser)+postID, userID).Val()
	curScore := pipeline.ZScore(redis.GetKey(redis.KeyPostVoteScore), postID).Val()
	curScore = (status-voteUserStatus)*score + curScore
	pipeline.ZIncrBy(redis.GetKey(redis.KeyPostVoteScore), curScore, postID)
	if status == 0 {
		pipeline.ZRem(redis.GetKey(redis.KeyPostVoteUser)+postID, userID)
	} else {
		pipeline.ZAdd(redis.GetKey(redis.KeyPostVoteUser)+postID, redis.MakeZ(status, userID))
	}
	_, err = pipeline.Exec()
	return
}
