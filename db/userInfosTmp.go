package db

import (
	"log"
	"time"
)

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
