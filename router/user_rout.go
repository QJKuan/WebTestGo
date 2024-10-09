package router

import (
	"WebTest/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterControl 注册
func RegisterControl(r *gin.Engine) {
	r.POST("/add", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")

		if db.RegisterDb(username, password) {
			c.String(http.StatusOK, "注册成功")
		} else {
			c.String(http.StatusInternalServerError, "注册失败")
		}
	})
}

// LoginControl 登录
func LoginControl(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if db.LoginDb(username, password) {
			c.String(http.StatusOK, "登录成功")
		} else {
			c.String(http.StatusOK, "登陆失败，用户名或密码错误")
		}
	})
}

// UpdateUser 更新用户信息
func UpdateUser(r *gin.Engine) {
	r.POST("/updateUser", func(c *gin.Context) {
		var user db.User
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println("更新User时，转换JSON异常")
		}
		fmt.Printf("%v", user)
	})
}
