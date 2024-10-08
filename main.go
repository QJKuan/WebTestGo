package main

import (
	"WebTest/config"
	"WebTest/db"
	"WebTest/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化配置文件
	config.CfgInit()
	//注册所有请求
	router.RequestInit(r)
	//注册数据库
	db.InitDb()

	//启动
	err := r.Run(config.SER_PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
