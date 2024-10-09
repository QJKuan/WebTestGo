package db

import (
	"errors"
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

// GetFilesInfos  获取所有可下载文件属性回写前端页面
/**
page 页数
pageSize 每页多少条
*/
func GetFilesInfos(page int, pageSize int) *[]FileInfo {
	//查询起始位置 即偏移量
	offSet := (page - 1) * pageSize

	var fis []FileInfo
	//分页查找
	result := DB.Offset(offSet).Limit(pageSize).Find(&fis)
	if result.Error != nil {
		fmt.Println(errors.New("分页查询异常"))
	}
	return &fis

}

func GetFilesInfo(fileName string) bool {
	fi := FileInfo{FileName: fileName}
	//分页查找
	result := DB.First(&fi)
	if result.Error != nil {
		fmt.Println(errors.New("查询异常"))
		return false
	}
	return result.RowsAffected > 0
}
