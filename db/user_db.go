package db

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Register struct {
	gorm.Model
	Username string
	Password string
}

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"Name"`
	Age          int       `json:"Age"`
	Gender       string    `json:"Gender"`
	Phone        string    `json:"Phone"`
	Address      string    `json:"Address"`
	Headportrait string    `json:"Headportrait"`
	Created      time.Time `gorm:"autoCreateTime"`
	Updated      time.Time `gorm:"autoUpdateTime"`
}

// RegisterDb 注册
func RegisterDb(username string, pwd string) bool {
	reg := Register{
		Username: username,
		Password: pwd,
	}
	tx := DB.Create(&reg)
	if tx.Error != nil {
		return false
	}
	fmt.Printf("用户 %v 插入数据成功，成功插入 %v 行 \n", reg.Username, tx.RowsAffected)
	return true
}

// LoginDb 登录
func LoginDb(username string, pwd string) bool {
	var reg Register
	tx := DB.First(&reg, "username=?", username)
	if tx.Error != nil {
		return false
	}
	if pwd == reg.Password {
		fmt.Printf("用户 %v 在 %v 登陆成功 \n", username, time.Now())
		return true
	}
	return false
}
