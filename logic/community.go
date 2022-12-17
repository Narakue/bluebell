package logic

import (
	"bluebell/dao"
	"bluebell/models"
)

func GetCommunityList() []*models.Community {
	communityList := dao.GetCommunityList()
	return communityList
}
