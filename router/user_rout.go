package router

import (
	"WebTest/db"
	"crypto/sha256"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"log"
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

		user := db.Register{
			Username: username,
			Password: password,
		}
		//去数据库确认用户密码是否正确
		res := db.LoginDb(user)
		if res == 3 {
			//设置token在cookie中
			userJ, err := sonic.Marshal(&user)
			if err != nil {
				c.String(http.StatusInternalServerError, "类型转换失败")
				log.Println(err.Error())
				return
			}
			//将参数转换为string
			token := fmt.Sprintf("%x", sha256.Sum256(userJ))
			//将Token存入TOKEN_MAP中
			TOKEN_MAP[token] = user
			//存入cookie
			c.SetCookie("token", token, 3600, "/", "localhost", false, true)
			c.String(http.StatusOK, "登录成功")
			return
		} else if res == 2 {
			c.String(http.StatusOK, "登陆失败，用户并未激活")
			return
		} else {
			c.String(http.StatusOK, "登陆失败，用户名或密码错误")
			return
		}
	})
}

// UpdateUser 更新用户信息
/*func UpdateUser(r *gin.Engine) {
	r.POST("/updateUser", func(c *gin.Context) {
		var user db.User
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println("更新User时，转换JSON异常")
		}
		fmt.Printf("%v", user)
	})
}*/

// SetUserInfoTmp 存储用户详细信息
func SetUserInfoTmp(r *gin.Engine) {
	r.POST("/setUserInfoTmp", func(c *gin.Context) {
		var uit db.UserInfosTmp
		err := c.BindJSON(&uit)
		if err != nil {
			c.String(http.StatusBadRequest, "参数异常")
			log.Println("Json数据转换异常")
			return
		}
		db.SetUserInfoTmp(uit, "zhangsan")

	})
}

// TokenVerify 登录校验逻辑
func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果请求的是静态资源路径，直接放行
		if c.Request.URL.Path == "/login" /*|| strings.HasSuffix(c.Request.URL.Path, ".css")*/ {
			// 放行
			c.Next()
			return
		}

		//校验token
		token, err := c.Cookie("token")
		if err == nil || token != "" {
			// 校验token值是否正确
			if _, exists := TOKEN_MAP[token]; exists {
				// 放行
				c.Next()
				return
			}
		}
		// 继续处理请求
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
