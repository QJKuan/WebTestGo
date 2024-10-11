package router

import (
	"WebTest/config"
	"WebTest/db"
	"crypto/sha256"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AdminLogin 管理员登录
func AdminLogin(r *gin.RouterGroup) {
	r.POST("/adminLogin", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("password")

		if username == config.ADM_UNE && pwd == config.ADM_PWD {
			admin := db.Register{
				Username: username,
				Password: pwd,
			}
			adminJ, err := sonic.Marshal(admin)
			if err != nil {
				log.Println(err.Error())
				c.String(http.StatusBadRequest, "请求参数异常")
			}
			token := fmt.Sprintf("%x", sha256.Sum256(adminJ))
			ADMIN_TOKEN[token] = admin
			c.String(http.StatusOK, "登录成功")
			return
		}

		c.String(http.StatusBadRequest, "用户名或密码错误")
	})
}
