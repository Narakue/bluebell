package logic

import (
	"bluebell/dao/redis"
	"bluebell/util"
	"strconv"
)

func RefreshToken(id int64) (aToken string, err error) {
	rdb := redis.GetRdb()
	aToken, err = util.RefreshToken(id)
	rdb.Set(strconv.FormatInt(id, 10)+util.AToken, aToken, util.ATokenExistTime)
	return
}
