package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

// Register 用户登录账户密码
type Register struct {
	ID        uint
	Username  string `gorm:"primarykey"`
	Password  string
	Able      int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// RegisterDb 注册
func RegisterDb(username string, pwd string) int {
	reg := Register{
		Username: username,
		Password: pwd,
		Able:     0,
	}
	result := DB.First(&reg)
	if result.Error != nil {
		//如果数据不存在则创建
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx := DB.Create(&reg)
			if tx.Error != nil {
				return 3
			}
			fmt.Printf("用户 %v 插入数据成功，成功插入 %v 行 \n", reg.Username, tx.RowsAffected)
			return 1
		}
		log.Println(result.Error)
		return 3
	}
	return 2
}

// LoginDb 登录
func LoginDb(user Register) int {
	//返回值介绍 :
	//1 : 用户名或密码错误 | 2 : 用户名未激活且未填写申请信息
	//3 : 用户未激活 | 4 : 没有此用户 | 5 : 登陆成功

	var reg Register
	//查询用户是否存在
	res := DB.First(&reg, "username=?", user.Username)

	//没有此用户
	if res.Error != nil {
		return 4
	}
	//查询用户是否激活
	if reg.Able == 0 {
		fmt.Printf("用户 %v 未激活且未填写申请信息 \n", user.Username)
		// 用户名未激活且未填写申请信息
		return 2
	}
	if reg.Able == 1 {
		fmt.Printf("用户 %v 未激活 \n", user.Username)
		// 用户名未激活
		return 3
	}
	if user.Password == reg.Password {
		fmt.Printf("用户 %v 在 %v 登陆成功 \n", user.Username, time.Now())
		// 登陆成功
		return 5
	}
	// 用户名或密码错误
	return 1
}

// UserActivate 用户激活成为最终用户 将 able 字段变为 2
func UserActivate(username string, txt *gorm.DB) bool {
	user := Register{Username: username}
	result := txt.Model(&user).Update("able", 2)
	if result.Error != nil {
		log.Printf("用户 %s 激活异常 : %s \n", username, result.Error.Error())
		return false
	}
	return true
}
