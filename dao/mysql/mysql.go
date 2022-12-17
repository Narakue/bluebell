package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

var sqlDB *gorm.DB

func Init() error {
	driverName := viper.GetString("mysql.driver_name")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(viper.GetInt("max_idle_conns"))
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(viper.GetInt("max_open_conns"))
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SingularTable(true) // 禁止表名为结构体的复数
	sqlDB = db
	return nil
}

func Close() {
	err := sqlDB.Close()
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return sqlDB
}
