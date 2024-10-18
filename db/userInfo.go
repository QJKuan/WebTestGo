package db

import (
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
