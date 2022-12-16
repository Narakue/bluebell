package dao

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func CheckoutUserIsExist(p *models.SignUpParam) bool {
	db := mysql.GetDB()
	user := new(models.User)
	result := db.Where("username = ?", p.Username).First(user)
	return result.RowsAffected > 0
}

func InsertUser(user *models.User) bool {
	db := mysql.GetDB()
	result := db.Create(user)
	return result.RowsAffected > 0
}

func Login(user *models.User) bool {
	db := mysql.GetDB()
	result := db.Where("username = ? and password = ?", user.Username, user.Password).First(user)
	return result.RowsAffected > 0
}
