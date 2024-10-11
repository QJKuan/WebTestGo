package router

import (
	"WebTest/config"
	"github.com/gin-gonic/gin"
)

func RequestInit(r *gin.Engine) {
	//全局相关请求
	{
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
		//初始化Filter需要的参数
		FilterInit()
		//登录
		LoginControl(r)
	}

	//用户相关请求
	{
		userR := r.Group("/user")
		//注册服务
		RegisterControl(userR)
		//上传用户信息
		//UpdateUser(userR)
		//用户详情临时表
		SetUserInfoTmp(userR)
	}

	//文件相关请求
	{
		r.MaxMultipartMemory = config.GBL_UPMEM << 20
		//上传文件
		UploadFile(r)
		//下载文件
		DownloadFile(r)
		//获取所有文件详细信息
		GetFileInfos(r)
	}

	//管理员相关请求
	{
		adminR := r.Group("/admin")
		AdminLogin(adminR)
	}

}
