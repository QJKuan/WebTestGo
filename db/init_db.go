package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 声明数据库连接为全局变量
var DB *gorm.DB

func InitDb() {
	//初始化数据库
	dns := "root:akn123@tcp(139.9.63.51:3522)/goTest?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常，终止程序")
	}
}
