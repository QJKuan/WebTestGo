package db

import (
	"WebTest/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 声明数据库连接为全局变量
var DB *gorm.DB
var TST *gorm.DB

func InitDb() {
	//初始化数据库
	dns := config.DB_DNS

	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常，终止程序")
	}

	//事务变量
	TST = DB.Begin()
	if TST.Error != nil {
		panic("事务开启失败")
	}
}
