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
	"regexp"
)

// USER_TOKEN 创建全局变量: 存储User Token的值
var USER_TOKEN map[string]interface{}
var ADMIN_TOKEN map[string]interface{}

var RE *regexp.Regexp

// FilterInit 初始化Filter需要的参数
func FilterInit() {
	//创建token的值存储map
	userMap := make(map[string]interface{})
	USER_TOKEN = userMap
	adminMap := make(map[string]interface{})
	ADMIN_TOKEN = adminMap

	RE = regexp.MustCompile(`/(.*?)/`)
}

// TokenVerify 登录校验逻辑
func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//添加正则对权限做隔离 有 admin 和 user 两个
		path := c.Request.URL.Path
		auth := RE.FindString(path)

		//登录注册页面直接放行
		if path == "/login" || path == "/user/add" {
			c.Next()
			return
		}

		//非登陆或注册页面需要获取cookie中的token
		token, err := c.Cookie("token")
		if err != nil {
			log.Println("Cookie存在问题 : " + err.Error())
			//如果报错或者为空，则直接拦截
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		//对于 user 用户的是否为登录状态的校验
		if auth == "/user/" {
			// 校验token值是否正确
			if _, exists := USER_TOKEN[token]; exists {
				// 放行
				c.Next()
				return
			}

			//未登录则拦截并返回登陆页面
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		//对于 admin 用户的是否为登录状态的校验
		if auth == "/admin/" {
			// 校验token值是否正确
			if _, exists := ADMIN_TOKEN[token]; exists {
				// 放行
				c.Next()
				return
			}

			//未登录则拦截并返回登陆页面
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// 剩余请求全部拦截
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
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

		//判断是否为 admin 的用户
		if username == config.ADM_UNE && password == config.ADM_PWD {
			adminJ, err := sonic.Marshal(user)
			if err != nil {
				log.Println(err.Error())
				c.String(http.StatusBadRequest, "请求参数异常")
				return
			}
			token := fmt.Sprintf("%x", sha256.Sum256(adminJ))
			ADMIN_TOKEN[token] = user
			c.SetCookie("token", token, 0, "/", "localhost", false, true)
			//fmt.Println("****Admin Token**** :", token)
			c.String(http.StatusOK, "管理员登录成功")
			return
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
			USER_TOKEN[token] = user
			//fmt.Println("****User Token**** :", token)
			//存入cookie
			c.SetCookie("token", token, 0, "/", "localhost", false, true)
			c.String(http.StatusOK, "用户登录成功")
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
