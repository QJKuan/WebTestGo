package db

import (
	"fmt"
	"time"
)

type FileInfo struct {
	FileName string `gorm:"primaryKey"`
	FileSize int64
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  time.Time `gorm:"autoUpdateTime"`
}

// SetFileInfo 上传文件时 设置文件属性
func SetFileInfo(fi FileInfo) {

	//查询安装包信息是否存在 存在则会获取到创建时间，不存在则会将当前时间变为创建时间 以便后续赋值
	oldFi := FileInfo{FileName: fi.FileName}
	err := DB.Find(&oldFi)
	if err.Error != nil {
		fmt.Println("Error:", err.Error)
		return
	}

	// 如果记录存在，则更新数据
	fi.Created = oldFi.Created
	DB.Save(&fi)
	fmt.Printf("更新文件 %v 成功 \n", fi.FileName)

}

// GetFilesInfo  获取所有可下载文件属性回写前端页面
func GetFilesInfo() {
	fi := FileInfo{}
	DB.First(&fi)
	fmt.Println(fi)
}
