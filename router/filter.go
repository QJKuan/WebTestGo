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
var USER_TOKEN map[string][]byte
var ADMIN_TOKEN map[string][]byte
var LINSHI_TOKEN map[string][]byte

var RE *regexp.Regexp

// FilterInit 初始化Filter需要的参数
func FilterInit() {
	//创建token的值存储map
	userMap := make(map[string][]byte)
	USER_TOKEN = userMap
	adminMap := make(map[string][]byte)
	ADMIN_TOKEN = adminMap
	linMap := make(map[string][]byte)
	LINSHI_TOKEN = linMap

	RE = regexp.MustCompile(`/(.*?)/`)
}

// TokenVerify 登录校验逻辑
func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//添加正则对权限做隔离 有 admin 和 user 两个
		path := c.Request.URL.Path
		auth := RE.FindString(path)

		//登录注册页面直接放行 退出登录也直接放行
		if path == "/login" || path == "/user/add" || path == "/logout" || path == "/user/setUserInfoTmp" {
			c.Next()
			return
		}

		//非登陆或注册页面需要获取cookie中的token
		token, err := c.Cookie("token")
		if err != nil {
			log.Println("Cookie存在问题 : " + err.Error())
			//如果报错或者为空，则直接拦截
			c.String(http.StatusBadRequest, "权限不足")
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
			c.String(http.StatusBadRequest, "权限不足")
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
			c.String(http.StatusBadRequest, "权限不足")
			c.Abort()
			return
		}

		// 剩余请求全部拦截
		c.String(http.StatusBadRequest, "权限不足")
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
			//保存 Token
			if !saveToken(user, c, "admin") {
				c.String(http.StatusBadRequest, "请求参数异常")
				return
			}
			c.String(http.StatusOK, "admin")
			return

		}

		//返回值介绍 :
		//1 : 用户名或密码错误 | 2 : 用户名未激活且未填写申请信息
		//3 : 用户未激活 | 4 : 没有此用户 | 5 : 登陆成功
		res := db.LoginDb(user)

		//登陆成功
		if res == 5 {
			//保存 Token
			if !saveToken(user, c, "user") {
				c.String(http.StatusBadRequest, "请求参数异常")
				return
			}
			c.String(http.StatusOK, "user")
			return
		} else if res == 2 {
			//用户名未激活且未填写申请信息
			//保存 Token
			if !saveToken(user, c, "linshi") {
				c.String(http.StatusBadRequest, "请求参数异常")
				return
			}
			c.String(http.StatusBadRequest, "ApplyFor")
			return
		} else if res == 3 {
			//用户未激活
			c.String(http.StatusBadRequest, "登陆失败，用户并未激活")
			return
		} else if res == 4 {
			c.String(http.StatusBadRequest, "登陆失败，用户名不存在")
			return
		} else {
			c.String(http.StatusBadRequest, "登陆失败，用户名或密码不正确")
			return
		}
	})
}

// Logout 退出登录
func Logout(r *gin.Engine) {
	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", 0, "/", "localhost", false, true)
		c.String(http.StatusOK, "退出登录成功")
		//c.Redirect(http.StatusFound, "/login")
		c.Abort()
	})
}

// 创建token并set进cookie中
func saveToken(user db.Register, c *gin.Context, name string) bool {
	userJson, err := sonic.Marshal(user)
	if err != nil {
		log.Println("user Json 转换异常 : " + err.Error())
		return false
	}

	//生成token
	token := fmt.Sprintf("%x", sha256.Sum256(userJson))
	if name == "admin" {
		ADMIN_TOKEN[token] = userJson
	} else if name == "user" {
		USER_TOKEN[token] = userJson
	} else {
		LINSHI_TOKEN[token] = userJson
	}
	c.SetCookie("token", token, 0, "/", "localhost", false, true)
	return true
}
