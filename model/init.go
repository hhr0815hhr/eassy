package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	dbRead  *gorm.DB
	dbWrite *gorm.DB
)

func InitDB(serverType string) {
	var (
		connStrW string
		connStrR string
		dbName   string
	)
	masterConf := viper.GetStringMapString("mysql.master.m1")
	slaverConf := viper.GetStringMapString("mysql.slaver")
	dbName = viper.GetString("dbName." + serverType)
	connStrW = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		masterConf["user"], //config.MysqlUser,
		masterConf["pass"], //config.MysqlPwd,
		masterConf["host"], //config.MysqlHost,
		masterConf["port"], //config.MysqlPort,
		dbName,             //config.MyDBName)
	)
	connStrR = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		slaverConf["user"], //config.MysqlUser,
		slaverConf["pass"], //config.MysqlPwd,
		slaverConf["host"], //config.MysqlHost,
		slaverConf["port"], //config.MysqlPort,
		dbName,             //config.MyDBName)
	)
	var err error
	dbRead, err = gorm.Open("mysql", connStrR)
	if err != nil {
		panic(err)
	}
	dbWrite, err = gorm.Open("mysql", connStrW)
	if err != nil {
		panic(err)
	}

	setDbPool()
}

//连接池
func setDbPool() {
	dbRead.DB().SetMaxIdleConns(10)
	dbRead.DB().SetMaxOpenConns(50)

	dbWrite.DB().SetMaxIdleConns(10)
	dbWrite.DB().SetMaxOpenConns(50)
}
