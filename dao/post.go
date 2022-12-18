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

func GetPostList() ([]*models.Post, error) {
	db := mysql.GetDB()
	postList := make([]*models.Post, 0)
	result := db.Find(&postList)
	return postList, result.Error
}
