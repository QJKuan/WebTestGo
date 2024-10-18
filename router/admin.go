package router

import (
	"WebTest/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserTmp(r *gin.RouterGroup) {
	r.POST("/activate", func(c *gin.Context) {
		username := c.PostForm("username")

		txt := db.DB.Begin()
		//添加用户表 删除临时表 激活注册表
		if !db.SetUserInfo(username, txt) || !db.DeleteUserInfoTmp(username, txt) || !db.UserActivate(username, txt) {
			c.String(http.StatusInternalServerError, "激活异常")
			txt.Rollback()
			return
		}
		txt.Commit()
		c.String(http.StatusOK, "激活成功")
	})
}
