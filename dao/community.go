package dao

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	db := mysql.GetDB()
	communityList := make([]*models.Community, 0)
	result := db.Find(&communityList)
	return communityList, result.Error
}

func GetCommunityDetailByID(id int64) (*models.Community, error) {
	db := mysql.GetDB()
	community := &models.Community{CommunityID: id}
	result := db.Where(community).First(community)
	return community, result.Error
}
