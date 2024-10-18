package db

import (
	"log"
	"time"
)

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

// SetUserInfo 插入最终用户信息
func SetUserInfo(ui UserInfo) bool {
	result := DB.Create(ui)
	if result.Error != nil {
		log.Printf("插入用户 %s 失败 报错 : %s \n", ui.Username, result.Error.Error())
		return false
	}
	return true
}

// GetUserInfos 分页查询最终用户表
func GetUserInfos(page int, pageSize int) *[]UserInfo {
	//查询起始位置 即偏移量
	offSet := (page - 1) * pageSize

	var ui []UserInfo
	//分页查找
	result := DB.Offset(offSet).Limit(pageSize).Find(&ui)
	if result.Error != nil {
		log.Printf("分页查询用表 UserInfo 数据异常 : %s \n", result.Error.Error())
	}
	return &ui
}
