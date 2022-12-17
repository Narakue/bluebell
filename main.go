package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/router"
	"bluebell/setting"
	"bluebell/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	gin.SetMode(gin.DebugMode)
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
	r := router.Init()
	err := r.Run(fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port")))
	if err != nil {
		fmt.Println("gin run fail", err)
	}
}
