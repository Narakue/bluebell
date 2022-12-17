package dao

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() []*models.Community {
	db := mysql.GetDB()
	communityList := make([]*models.Community, 0)
	db.Find(&communityList)
	return communityList
}
