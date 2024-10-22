package template

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func ExcelTem() {
	//创建 excel 文件
	excelFile := excelize.NewFile()
	defer func() {
		if err := excelFile.Close(); err != nil {
			log.Println("关闭 excel 文件异常 : " + err.Error())
		}
	}()

	//创建一个工作表
	sheetName := "用户详情"
	sheet1, err := excelFile.NewSheet(sheetName)
	if err != nil {
		log.Println("创建 excel 工作表异常 : " + err.Error())
		return
	}

	//设置默认工作表
	excelFile.SetActiveSheet(sheet1)

	// 单元格的值
	tem := []struct {
		line  string
		value string
	}{
		{"A1", "用户名"},
		{"B1", "区域"},
		{"C1", "建设单位"},
		{"F1", "集成单位"},
		{"I1", "开发单位"},
		{"L1", "应用名称"},
		{"M1", "部署环境"},
		{"N1", "中间件名称"},
		{"O1", "创建时间"},
		{"C2", "单位名称"},
		{"D2", "联系人"},
		{"E2", "手机号"},
		{"F2", "单位名称"},
		{"G2", "联系人"},
		{"H2", "手机号"},
		{"I2", "单位名称"},
		{"J2", "联系人"},
		{"K2", "手机号"},
	}
	//设置单元格的值
	for _, v := range tem {
		err2 := excelFile.SetCellValue(sheetName, v.line, v.value)
		if err2 != nil {
			log.Printf("设置单元格 %s 值异常 : %s\n", v.line, err2.Error())
			return
		}
	}

	//需要合并的单元格
	mergeCells := []struct {
		start string
		end   string
	}{
		{"A1", "A2"},
		{"B1", "B2"},
		{"L1", "L2"},
		{"M1", "M2"},
		{"N1", "N2"},
		{"O1", "O2"},
		{"C1", "E1"},
		{"F1", "H1"},
		{"I1", "K1"},
	}
	//合并单元格
	for _, v := range mergeCells {
		err2 := excelFile.MergeCell(sheetName, v.start, v.end)
		if err2 != nil {
			log.Printf("合并单元格 %s 到 %s 异常 : %s\n", v.start, v.end, err2.Error())
			return
		}
	}

	//单元格格式
	style, err3 := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Style: 1, Color: "#000000"},
			{Type: "right", Style: 1, Color: "#000000"},
			{Type: "top", Style: 1, Color: "#000000"},
			{Type: "bottom", Style: 1, Color: "#000000"},
		},
	})
	if err3 != nil {
		log.Printf("excel样式生成异常 : %s\n", err3.Error())
		return
	}
	//设置单元格格式
	for i := 1; i < 100; i++ {
		num := strconv.Itoa(i)
		err4 := excelFile.SetCellStyle(sheetName, "A"+num, "O"+num, style)
		if err4 != nil {
			log.Printf("excel样式使用异常 : %s\n", err4.Error())
			return
		}
	}

	// 保存Excel文件
	if err5 := excelFile.SaveAs("./file/用户详情.xlsx"); err5 != nil {
		log.Println("保存 excel 文件异常 : " + err5.Error())
		return
	}
	log.Println("Excel文件已成功创建并保存为 用户详情.xlsx")
}
