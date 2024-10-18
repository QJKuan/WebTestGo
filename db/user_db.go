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
	Username      string    `gorm:"primaryKey"`
	IP            string    `json:"IP" form:"IP"`
	Area          string    `json:"Area" form:"Area"`
	JsName        string    `json:"JsName" form:"JsName"`
	JsLinkman     string    `json:"JsLinkman" form:"JsLinkman"`
	JsPhone       string    `json:"JsPhone" form:"JsPhone"`
	JcName        string    `json:"JcName" form:"JcName"`
	JcLinkman     string    `json:"JcLinkman" form:"JcLinkman"`
	JcPhone       string    `json:"JcPhone" form:"JcPhone"`
	KfName        string    `json:"KfName" form:"KfName"`
	KfLinkman     string    `json:"KfLinkman" form:"KfLinkman"`
	KfPhone       string    `json:"KfPhone" form:"KfPhone"`
	AppName       string    `json:"AppName" form:"AppName"`
	Env           string    `json:"Env" form:"Env"`
	MiddleProduct string    `json:"MiddleProduct" form:"MiddleProduct"`
	Created       time.Time `gorm:"autoCreateTime"`
	Updated       time.Time `gorm:"autoUpdateTime"`
}

// UserInfosTmp 用户信息临时表
type UserInfosTmp struct {
	Username      string    `gorm:"primaryKey"`
	IP            string    `json:"IP" form:"IP"`
	Area          string    `json:"Area" form:"Area"`
	JsName        string    `json:"JsName" form:"JsName"`
	JsLinkman     string    `json:"JsLinkman" form:"JsLinkman"`
	JsPhone       string    `json:"JsPhone" form:"JsPhone"`
	JcName        string    `json:"JcName" form:"JcName"`
	JcLinkman     string    `json:"JcLinkman" form:"JcLinkman"`
	JcPhone       string    `json:"JcPhone" form:"JcPhone"`
	KfName        string    `json:"KfName" form:"KfName"`
	KfLinkman     string    `json:"KfLinkman" form:"KfLinkman"`
	KfPhone       string    `json:"KfPhone" form:"KfPhone"`
	AppName       string    `json:"AppName" form:"AppName"`
	Env           string    `json:"Env" form:"Env"`
	MiddleProduct string    `json:"MiddleProduct" form:"MiddleProduct"`
	Created       time.Time `gorm:"autoCreateTime"`
	Updated       time.Time `gorm:"autoUpdateTime"`
}

/*type UserInfosTmp struct {
	ID      uint        `gorm:"primaryKey"`
	IP      string      `form:"IP"`
	Area    string      `form:"Area"`
	JianShe CompanyInfo `form:"JianShe"`
	JiCheng CompanyInfo `form:"JiCheng"`
	KaiFa   CompanyInfo `form:"KaiFa"`
	AppName string      `form:"AppName"`
	Env     string      `form:"Env"`
	Created time.Time   `gorm:"autoCreateTime"`
	Updated time.Time   `gorm:"autoUpdateTime"`
}
type CompanyInfo struct {
	Name    string `json:"Name" form:"Name"`
	Linkman string `json:"Linkman" form:"Linkman"`
	Phone   string `json:"Phone" form:"Linkman"`
}*/

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

// SetUserInfoTmp 插入临时用户信息
func SetUserInfoTmp(uit UserInfosTmp) bool {
	//开始事务
	ts := DB.Begin()

	//插入申请数据
	resultU := ts.Create(&uit)
	if resultU.Error != nil {
		log.Println("插入数据异常 : " + resultU.Error.Error())
		ts.Rollback()
		return false
	}

	//将 user 的 able 更新为 1
	resultR := ts.Model(&Register{Username: uit.Username}).Update("able", 1)
	if resultR.Error != nil {
		log.Println("更新数据异常 : " + resultU.Error.Error())
		ts.Rollback()
		return false
	}
	ts.Commit()
	return true
}
