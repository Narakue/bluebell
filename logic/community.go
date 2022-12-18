package logic

import (
	"bluebell/dao"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	communityList, err := dao.GetCommunityList()
	return communityList, err
}

func GetCommunityDetailByID(id int64) (*models.Community, error) {
	community, err := dao.GetCommunityDetailByID(id)
	return community, err
}
