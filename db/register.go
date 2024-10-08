package db

import (
	"gorm.io/gorm"
	"time"
)

type Registerss struct {
	gorm.Model
	Username string
	Password string
}

type Userss struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Age          int
	Gender       string
	Phone        string
	Address      string
	Headportrait string
	Created      time.Time `gorm:"autoCreateTime"`
	Updated      time.Time `gorm:"autoCreateTime"`
}

func Login() {
	/*	dsn := "root:akn123@tcp(139.9.63.51:3522)/goTest?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("数据库连接异常，终止程序")
		}
		var re Register
		db.First(&re, 1)
		fmt.Println(re)

		user := User{
			ID:           1,
			Name:         "zhangsan",
			Age:          18,
			Gender:       "男",
			Phone:        "13926557890",
			Address:      "广东省深圳市",
			Headportrait: "http://qwe.com",
		}
		result := db.Create(&user)
		fmt.Println(result)*/

}
