package router

import (
	"WebTest/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CheckUserTmp 激活用户
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

// GetUserInfos 获取所有最终用户
func GetUserInfos(r *gin.RouterGroup) {
	r.POST("/getUserInfo", func(c *gin.Context) {
		//获取分页参数
		pag := Pagination{}
		err := c.Bind(&pag)
		if err != nil {
			c.String(http.StatusBadRequest, "参数异常,无法转换为Json")
			log.Println(err.Error())
			return
		}
		if pag.PageSize >= 20 {
			pag.PageSize = 20
		}

		infos, pagin := db.GetUserInfos(pag.Page, pag.PageSize)
		c.JSON(http.StatusOK, gin.H{
			"infos": infos,
			"pagin": pagin,
		})
	})
}

// GetUserInfosTmp 获取所有临时用户
func GetUserInfosTmp(r *gin.RouterGroup) {
	r.POST("/getUserTmp", func(c *gin.Context) {
		//获取分页参数
		pag := Pagination{}
		err := c.Bind(&pag)
		if err != nil {
			c.String(http.StatusBadRequest, "参数异常,无法转换为Json")
			log.Println(err.Error())
			return
		}
		if pag.PageSize >= 20 {
			pag.PageSize = 20
		}

		tmp, pagin := db.GetUserInfosTmp(pag.Page, pag.PageSize)
		c.JSON(http.StatusOK, gin.H{
			"tmp":   tmp,
			"pagin": pagin,
		})
	})
}

// DeleteUserTmp 临时表单不合格，删除临时表
func DeleteUserTmp(r *gin.RouterGroup) {
	r.POST("/deleteTmp", func(c *gin.Context) {
		username := c.PostForm("username")

		txt := db.DB.Begin()
		//添加用户表 删除临时表 激活注册表
		if !db.DeleteUserInfoTmp(username, txt) || !db.UserSendBack(username, txt) {
			c.String(http.StatusInternalServerError, "拒绝异常")
			txt.Rollback()
			return
		}
		txt.Commit()
		c.String(http.StatusOK, "拒绝成功")

	})
}
