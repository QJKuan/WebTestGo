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
	ID        uint   /*`gorm:"primarykey"`*/
	Username  string `gorm:"primarykey"`
	Password  string
	Able      int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

/*type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `json:"Name"`
	Age          int       `json:"Age"`
	Gender       string    `json:"Gender"`
	Phone        string    `json:"Phone"`
	Address      string    `json:"Address"`
	Headportrait string    `json:"Headportrait"`
	Created      time.Time `gorm:"autoCreateTime"`
	Updated      time.Time `gorm:"autoUpdateTime"`
}*/

// UserInfo 用户信息详情表
type UserInfo struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	IP        string    `json:"IP"`
	Area      int       `json:"Area"`
	JsName    string    `json:"JsName"`
	JsLinkman string    `json:"JsLinkman"`
	JsPhone   string    `json:"JsPhone"`
	JcName    string    `json:"JcName"`
	JcLinkman string    `json:"JcLinkman"`
	JcPhone   string    `json:"JcPhone"`
	KfName    string    `json:"KfName"`
	KfLinkman string    `json:"KfLinkman"`
	KfPhone   string    `json:"KfPhone"`
	AppName   string    `json:"AppName"`
	Env       string    `json:"Env"`
	Created   time.Time `gorm:"autoCreateTime"`
	Updated   time.Time `gorm:"autoUpdateTime"`
}

// UserInfosTmp 用户信息临时表
type UserInfosTmp struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	IP        string    `json:"IP"`
	Area      int       `json:"Area"`
	JsName    string    `json:"JsName"`
	JsLinkman string    `json:"JsLinkman"`
	JsPhone   string    `json:"JsPhone"`
	JcName    string    `json:"JcName"`
	JcLinkman string    `json:"JcLinkman"`
	JcPhone   string    `json:"JcPhone"`
	KfName    string    `json:"KfName"`
	KfLinkman string    `json:"KfLinkman"`
	KfPhone   string    `json:"KfPhone"`
	AppName   string    `json:"AppName"`
	Env       string    `json:"Env"`
	Created   time.Time `gorm:"autoCreateTime"`
	Updated   time.Time `gorm:"autoUpdateTime"`
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
	var reg Register
	//查询用户是否存在
	DB.First(&reg, "username=?", user.Username)

	//查询用户时候激活
	if reg.Able == 0 {
		fmt.Printf("用户 %v 未激活 \n", user.Username)
		// 返回 2 为用户名未激活
		return 2
	}
	if user.Password == reg.Password {
		fmt.Printf("用户 %v 在 %v 登陆成功 \n", user.Username, time.Now())
		// 返回 2 为登陆成功
		return 3
	}
	// 返回 1 为用户名或密码错误
	return 1
}

// SetUserInfoTmp 插入临时用户信息
func SetUserInfoTmp(uit UserInfosTmp, username string) {
	var user Register
	//查询用户id
	DB.First(&user, "username=?", username)

	uit.ID = user.ID
	DB.Create(&uit)
}
