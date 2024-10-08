package router

import (
	"WebTest/config"
	"WebTest/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
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

// UploadFile 上传文件
func UploadFile(r *gin.Engine) {
	r.MaxMultipartMemory = config.GBL_UPMEM << 20
	r.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		fmt.Printf("文件名称： %v \n大小：     %v MB \n", file.Filename, file.Size>>20)

		//判断是否包含尾缀 "."
		if !strings.Contains(file.Filename, ".") {
			c.String(http.StatusBadRequest, "上传的文件格式不正确，不包含尾缀")
			return

		}

		//获取文件尾缀
		re := regexp.MustCompile(`[^.]*$`)
		suf := re.FindString(file.Filename)

		if suf == "zip" || suf == "7z" || suf == "gz" {
			log.Printf("文件名称 ： %v \n", file.Filename)

			//上传文件名称
			dst := "./file/" + file.Filename

			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				log.Println("保存文件异常")
				c.String(http.StatusInternalServerError, "文件保存异常")
				return
			}

			c.String(http.StatusOK, "文件上传成功")
			return
		}

		c.String(http.StatusBadRequest, "上传的文件格式不正确，格式只能为 zip , 7z , gz")
	})
}
