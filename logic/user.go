package logic

import (
	"bluebell/dao"
	"bluebell/models"
	"bluebell/util"
	"errors"
)

func SignUp(p *models.SignUpParam) (err error) {
	var res bool
	res = dao.SignUp(p)
	if res {
		return errors.New("user already exist")
	}
	uid := util.GenID()
	user := models.User{UserID: uid, Username: p.Username, Password: p.Password}
	res = dao.InsertUser(&user)
	if !res {
		return errors.New("user create fail")
	}
	return err
}
