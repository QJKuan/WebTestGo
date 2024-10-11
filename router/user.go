package router

import (
	"WebTest/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// RegisterControl 注册
func RegisterControl(r *gin.RouterGroup) {
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

// UpdateUser 更新用户信息
/*func UpdateUser(r *gin.RouterGroup) {
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
func SetUserInfoTmp(r *gin.RouterGroup) {
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
