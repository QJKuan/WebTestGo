package router

import (
	"WebTest/config"
	"WebTest/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// USER_TOKEN 存储User Token的值
var USER_TOKEN map[string]string
var ADMIN_TOKEN map[string]string
var LINSHI_TOKEN map[string]string

// TOKEN_EXIST 判断登录状态
var TOKEN_EXIST map[string]string

var RE *regexp.Regexp
var FILE_RE *regexp.Regexp

// FilterInit 初始化Filter需要的参数
func FilterInit() {
	//创建token的值存储map
	USER_TOKEN = make(map[string]string)
	ADMIN_TOKEN = make(map[string]string)
	LINSHI_TOKEN = make(map[string]string)

	//存入用户名和token 查看此用户是否处于登录状态
	TOKEN_EXIST = make(map[string]string)

	RE = regexp.MustCompile(`/(.*?)/`)
	FILE_RE = regexp.MustCompile(`[^.]*$`)
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
		c.SetCookie("token", "", 0, "/", config.COOKIE_HOST, false, true)
		c.String(http.StatusOK, "退出登录成功")
		//c.Redirect(http.StatusFound, "/login")
		c.Abort()
	})
}

// 创建token并set进cookie中
func saveToken(user db.Register, c *gin.Context, name string) bool {
	//判断此用户是否已经登录 已经登录则直接拿之前的token返回 未登录则直接生成token
	username := user.Username
	tokenExist, exist := TOKEN_EXIST[username]
	if exist {
		c.SetCookie("token", tokenExist, 0, "/", config.COOKIE_HOST, false, true)
		return true
	}

	//生成token
	uid, errRan := uuid.NewRandom()
	if errRan != nil {
		log.Println("随机数创建失败 : " + errRan.Error())
		return false
	}
	token := strings.ReplaceAll(uid.String(), "-", "")

	//token := fmt.Sprintf("%x", sha256.Sum256(userJson))

	//将 token 存入 Map 中
	if name == "admin" {
		ADMIN_TOKEN[token] = username
	} else if name == "user" {
		USER_TOKEN[token] = username
	} else {
		LINSHI_TOKEN[token] = username
	}

	TOKEN_EXIST[username] = token
	c.SetCookie("token", token, 0, "/", config.COOKIE_HOST, false, true)
	return true
}

// Timer 定时任务清空map
func Timer() {
	go func() {
		tick := time.NewTicker(24 * 7 * time.Hour)
		defer tick.Stop() // 确保在协程结束时停止定时器
		for {
			select {
			case <-tick.C:
				FilterInit()
				fmt.Println("===============清空所有登录信息=================")
			}
		}
	}()
}
