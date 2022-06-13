package models

import (
	"github.com/solost23/tools/logger"
	"github.com/solost23/tools/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

func NewMysqlConnect() (connect *gorm.DB) {
	var err error
	connect, err = mysql.NewMysqlConnect(&mysql.Config{
		UserName:        viper.GetString("connections.mysql.video_server.user"),
		Password:        viper.GetString("connections.mysql.video_server.password"),
		Host:            viper.GetString("connections.mysql.video_server.host"),
		Port:            viper.GetInt("connections.mysql.video_server.port"),
		DB:              viper.GetString("connections.mysql.video_server.db"),
		Charset:         viper.GetString("connections.mysql.video_server.charset"),
		MaxIdleConn:     10,
		MaxOpenConn:     100,
		ConnMaxLifeTime: time.Hour,
		Logger:          logger.NewMysqlLogger(),
	})
	if err != nil {
		panic(err)
	}
	return connect
}
