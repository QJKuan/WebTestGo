package router

import "github.com/gin-gonic/gin"

func CheckUserTmp(r *gin.RouterGroup) {
	r.POST("/activate", func(c *gin.Context) {
		//username := c.PostForm("username")

	})
}
