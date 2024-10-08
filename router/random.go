package router

import (
	"WebTest/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func GetRandom(route *gin.Engine) {

	//route.LoadHTMLFiles("./goindex.html")
	//route.GET("/xx", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "goindex.html", nil)
	//})

	route.MaxMultipartMemory = 300 << 20
	route.POST("/getRan", func(c *gin.Context) {

		//version : "9" 为v9 | "10" 为v10 | "10a" 为10敏捷版 | "10b" 为10内嵌版 | "11a" 为v11敏捷版 | "11b" 为内嵌版
		// "12a" 为 ALB_X86 | "12b" 为 ARM | "13" 为 ADMQ | "14a" 为 AMDC_X86 | "14b" 为 AMDC_ARM
		ver := c.PostForm("version")
		uid, err := uuid.NewRandom()
		if err != nil {
			c.String(http.StatusInternalServerError, "随机数创建失败")
			return
		}
		random := strings.ReplaceAll(uid.String(), "-", "")
		//存入Map
		RAN_MAP[random] = ver

		url := config.WRK_URL + random
		c.String(http.StatusOK, url)
	})

	route.GET("/download/:rand", func(c *gin.Context) {
		rand := c.Param("rand")
		ver, exist := RAN_MAP[rand]
		if !exist {
			c.String(http.StatusBadRequest, "不正确的url地址")
			return
		}

		defer delete(RAN_MAP, rand)
		switch ver {
		case "9":
			c.FileAttachment("./file/AAS-V9.zip", "AAS-V9.zip")
		case "10":
			c.FileAttachment("./file/AAS-V10.zip", "AAS-V10.zip")
		case "10a":
			c.FileAttachment("./file/AAMS-V10.zip", "AAMS-V10.zip")
		case "10b":
			c.FileAttachment("./file/AAMS-V10-embed.zip", "AAMS-V10-embed.zip")
		//case "11a":
		//	c.FileAttachment("./file/AAMS-V11.zip", "AAMS-V11.zip")
		//case "11b":
		//	c.FileAttachment("AAMS-V11-embed.zip", "AAMS-V11-embed.zip")
		case "12a":
			c.FileAttachment("./file/ALB_X86.zip", "ALB_X86.zip")
		case "12b":
			c.FileAttachment("./file/ALB_ARM.zip", "ALB_ARM.zip")
		case "13":
			c.FileAttachment("./file/ADMQ.zip", "ADMQ.zip")
		case "14a":
			c.FileAttachment("./file/AMDC_X86.zip", "AMDC_X86.zip")
		case "14b":
			c.FileAttachment("./file/AMDC_ARM.zip", "AMDC_ARM.zip")

		default:
			c.String(http.StatusBadRequest, "不正确的参数")
		}
	})
}

func Timer() {
	go func() {
		tick := time.NewTicker(12 * time.Hour)
		defer tick.Stop() // 确保在协程结束时停止定时器
		for {
			select {
			case <-tick.C:
				RAN_MAP = make(map[string]string)
				fmt.Println("===============清空链接=================")
			}
		}
	}()
}
