package dao

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func CreatePost(post *models.Post) (bool, error) {
	db := mysql.GetDB()
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
