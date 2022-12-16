package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routers"
	"bluebell/setting"
	"bluebell/util"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := setting.Init(); err != nil {
		fmt.Println("init setting fail", err)
	}
	if err := logger.Init(); err != nil {
		fmt.Println("init logger fail", err)
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {

		}
	}(zap.L())
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql fail", err)
	}
	defer mysql.Close()
	if err := redis.Init(); err != nil {
		fmt.Println("init redis fail", err)
	}
	defer redis.Close()
	if err := util.Init(viper.GetString("start_time"), viper.GetInt64("machine_id")); err != nil {
		fmt.Println("init snowflake fail", err)
	}
	if err := routers.Init(); err != nil {
		fmt.Println("init routers fail", err)
	}
}
