package router

import (
	"WebTest/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// UploadFile 上传文件
func UploadFile(r *gin.RouterGroup) {
	r.POST("/upload", func(c *gin.Context) {

		file, _ := c.FormFile("file")
		fmt.Printf("文件名称： %v \n大小：     %v MB \n", file.Filename, file.Size>>20)

		//判断是否包含尾缀 "."
		if !strings.Contains(file.Filename, ".") {
			c.String(http.StatusBadRequest, "上传的文件格式不正确，不包含尾缀")
			return
		}

		//获取文件尾缀
		suf := FILE_RE.FindString(file.Filename)

		//查看文件尾缀是否正确
		if suf == "zip" || suf == "7z" || suf == "gz" {
			log.Printf("文件名称 ： %v \n", file.Filename)

			//上传文件名称
			dst := "./file/" + file.Filename

			//将文件信息传入数据库留存
			db.SetFileInfo(db.FileInfo{
				FileName: file.Filename,
				FileSize: file.Size,
			})

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

// DownloadFile 文件下载功能
func DownloadFile(r *gin.RouterGroup) {
	r.POST("/downfile", func(c *gin.Context) {
		fileName := c.PostForm("fileName")
		//判断是否有此文件名称
		if !db.GetFilesInfo(fileName) {
			c.String(http.StatusBadRequest, "文件名称有误")
			return
		}
		//返回文件
		c.FileAttachment("./file/"+fileName, fileName)
	})
}

// Pagination 接受分页参数的结构体
type Pagination struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

// GetFileInfos 获取文件详细信息
func GetFileInfos(r *gin.RouterGroup) {
	r.POST("/getFileInfos", func(c *gin.Context) {
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

		infos, pagin := db.GetFilesInfos(pag.Page, pag.PageSize)
		c.JSON(http.StatusOK, gin.H{
			"files": infos,
			"pagin": pagin,
		})
	})
}

// DeleteFile 文件删除功能
func DeleteFile(r *gin.RouterGroup) {
	r.POST("/deleteFile", func(c *gin.Context) {
		fileName := c.PostForm("fileName")

		//判断是否有此文件名称
		if !db.GetFilesInfo(fileName) {
			c.String(http.StatusBadRequest, "文件名称有误")
			return
		}

		//开启事务删除数据库记录并删除文件
		if !db.DeleteFilesInfo(fileName) {
			c.String(http.StatusBadRequest, "删除记录异常")
			return
		}

		c.String(http.StatusOK, "文件删除成功")
	})
}
