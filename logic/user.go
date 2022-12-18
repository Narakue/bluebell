package logic

import (
	"bluebell/dao"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/util"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
)

func SignUp(p *models.SignUpParam) error {
	res, err := dao.CheckoutUserIsExist(p)
	if err != nil {
		return err
	}
	if res {
		return errors.New("user already exist")
	}
	uid := util.GenID()
	user := &models.User{UserID: uid, Username: p.Username, Password: encryptPassword(p.Password)}
	res, err = dao.InsertUser(user)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("user create fail")
	}
	return nil
}

func Login(p *models.LoginParam) (aToken string, rToken string, err error) {
	user := &models.User{Username: p.Username, Password: p.Password}
	user.Password = encryptPassword(user.Password)
	res, err := dao.Login(user)
	if !res {
		err = errors.New("login fail")
		return
	}
	aToken, rToken, err = util.GenerateToken(user.UserID, user.Password)
	rdb := redis.GetRdb()
	rdb.Set(strconv.FormatInt(user.UserID, 10)+util.AToken, aToken, util.ATokenExistTime)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte("zyz"))
	sum := h.Sum([]byte(password))
	return hex.EncodeToString(sum)
}

func GetUserByID(id int64) (*models.User, error) {
	return dao.GetUserByID(id)
}
