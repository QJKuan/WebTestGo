package server

import (
	"WebTest/db"
	"bytes"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

// 方便 setCellValue 方法取值
var excelFile *excelize.File

func CreateExcel() *bytes.Buffer {
	//打开模板表
	var err1 error
	excelFile, err1 = excelize.OpenFile("./template/UserInfoTem.xlsx")
	if err1 != nil {
		log.Printf("打开 excel 工作表 UserInfoTem.xlsx 异常 : " + err1.Error())
		return nil
	}
	defer func() {
		if err := excelFile.Close(); err != nil {
			log.Println("关闭 excel 文件异常 : " + err.Error())
		}
	}()
	//获取表名称
	sheetName := excelFile.WorkBook.Sheets.Sheet[0].Name

	//获取所有数据库的最终用户信息
	userInfos := db.GetAllUserInfo()
	//行数起始值
	i := 3

	//插入值到 excel 中
	for _, ui := range userInfos {
		num := strconv.Itoa(i)
		setCellValue(sheetName, "A"+num, ui.Username)
		setCellValue(sheetName, "B"+num, ui.Area)
		setCellValue(sheetName, "C"+num, ui.JsName)
		setCellValue(sheetName, "D"+num, ui.JsLinkman)
		setCellValue(sheetName, "E"+num, ui.JsPhone)
		setCellValue(sheetName, "F"+num, ui.JcName)
		setCellValue(sheetName, "G"+num, ui.JcLinkman)
		setCellValue(sheetName, "H"+num, ui.JcPhone)
		setCellValue(sheetName, "I"+num, ui.KfName)
		setCellValue(sheetName, "J"+num, ui.KfLinkman)
		setCellValue(sheetName, "K"+num, ui.KfPhone)
		setCellValue(sheetName, "L"+num, ui.AppName)
		setCellValue(sheetName, "M"+num, ui.Env)
		setCellValue(sheetName, "N"+num, ui.MiddleProduct)
		setCellValue(sheetName, "O"+num, ui.Created)
		i++
	}

	fileStream, err2 := excelFile.WriteToBuffer()
	if err2 != nil {
		log.Println("将 excel 文件转换为 buffer 流异常 : " + err2.Error())
		return nil
	}
	return fileStream
	/*if err5 := excelFile.SaveAs("./file/用户详情.xlsx"); err5 != nil {
		log.Println("保存 excel 文件异常 : " + err5.Error())
		return
	}
	log.Println("Excel文件已成功创建并保存为 用户详情.xlsx")*/
}

// setCellValue 设置单元格值 主要是为了通一错误处理
func setCellValue(sheetName, cell string, value interface{}) {
	if err := excelFile.SetCellValue(sheetName, cell, value); err != nil {
		log.Printf("设置单元格 %s 值异常: %s", cell, err.Error())
	}
}

// incrementChar 将字符从 A 到 O 进行增量取值
func incrementChar(c byte) byte {
	if c >= 'A' && c < 'O' {
		return c + 1
	}
	return 'A' // 如果超出 O，则返回 A
}
