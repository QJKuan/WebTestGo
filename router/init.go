package router

import (
	"WebTest/config"
	"github.com/gin-gonic/gin"
)

func RequestInit(r *gin.Engine) {
	//全局相关请求
	{
		//设置文件上传下载一次可用内存
		r.MaxMultipartMemory = config.GBL_UPMEM << 20
		//取消跨域拦截
		r.Use(func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
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
		//退出登录
		Logout(r)
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
		//下载文件
		DownloadFile(userR)
		//获取所有文件详细信息
		GetFileInfos(userR)
	}

	//文件相关请求
	{

	}

	//管理员相关请求
	{
		adminR := r.Group("/admin")
		//上传文件
		UploadFile(adminR)
		//下载文件
		DownloadFile(adminR)
		//获取所有文件详细信息
		GetFileInfos(adminR)
		//文件删除
		DeleteFile(adminR)
		//激活临时用户
		CheckUserTmp(adminR)
		//获取所有最终用户
		GetUserInfos(adminR)
		//获取所有临时用户
		GetUserInfosTmp(adminR)
		//删除临时表 重新填写
		DeleteUserTmp(adminR)
	}

}
