package router

import (
	"WebTest/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminLogin 管理员登录
func AdminLogin(r *gin.RouterGroup) {
	r.POST("/adminLogin", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("password")

		if username == config.ADM_UNE && pwd == config.ADM_PWD {
			c.String(http.StatusOK, "登录成功")
			return
		}

		c.String(http.StatusBadRequest, "用户名或密码错误")
	})
}
