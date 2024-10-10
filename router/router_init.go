package router

import "github.com/gin-gonic/gin"

func RequestInit(r *gin.Engine) {
	//取消跨域拦截
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//拦截未登录请求
	r.Use(TokenVerify())

	//注册服务
	RegisterControl(r)
	//登录
	LoginControl(r)
	//上传用户信息
	//UpdateUser(r)
	//用户详情临时表
	SetUserInfoTmp(r)

	//上传文件
	UploadFile(r)
	//下载文件
	DownloadFile(r)
	//获取所有文件详细信息
	GetFileInfos(r)

	//获取下载地址以及下载aas相关包
	CreateMap()
	GetRandom(r)
	Timer()

}

// RAN_MAP 创建全局变量: 存储aas随机数和版本号
var RAN_MAP map[string]string

// TOKEN_MAP 创建全局变量: 存储Token的值
var TOKEN_MAP map[string]interface{}

func CreateMap() {
	//创建随机数存储的map
	ranMap := make(map[string]string)
	RAN_MAP = ranMap
	//创建token的值存储map
	tokMap := make(map[string]interface{})
	TOKEN_MAP = tokMap
}
