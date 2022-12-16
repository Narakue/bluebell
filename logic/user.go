package logic

import (
	"bluebell/dao"
	"bluebell/models"
	"bluebell/util"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func SignUp(p *models.SignUpParam) error {
	res := dao.CheckoutUserIsExist(p)
	if res {
		return errors.New("user already exist")
	}
	uid := util.GenID()
	user := &models.User{UserID: uid, Username: p.Username, Password: encryptPassword(p.Password)}
	res = dao.InsertUser(user)
	if !res {
		return errors.New("user create fail")
	}
	return nil
}

func Login(p *models.LoginParam) error {
	user := &models.User{Username: p.Username, Password: p.Password}
	user.Password = encryptPassword(user.Password)
	res := dao.Login(user)
	if !res {
		return errors.New("login fail")
	}
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte("zyz"))
	sum := h.Sum([]byte(password))
	return hex.EncodeToString(sum)
}
