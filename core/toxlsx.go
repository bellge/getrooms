package core

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/bellge/getrooms/tv"
	"github.com/tealeg/xlsx"
)

func initxlsx(filename string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	strs := []string{"时间", "触手TV", "飞云TV", "大神TV", "企鹅电竞"}
	for _, v := range strs {
		cell := row.AddCell()
		cell.Value = v
	}
	file.Save(filename)
	tv.First = false
}
func W2xls() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//	var cell *xlsx.Cell
	var err error
	filename := "data2.xlsx"
	if _, err := os.Stat(filename); err != nil {
		initxlsx(filename)
		//		return
	}
	file, _ = xlsx.OpenFile(filename)

	//	file = xlsx.NewFile()
	sheet = file.Sheet["Sheet1"]
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	//	cell = row.AddCell()

	//	cell.Value = "时间"
	//创建表头
	//	strs := []string{"时间", "触手TV", "飞云TV", "大神TV", "企鹅电竞"}
	//	if first {
	//		for _, v := range strs {
	//			cell := row.AddCell()
	//			cell.Value = v
	//		}
	//		file.Save("data2.xlsx")
	//		return
	//	}
	cell := row.AddCell()
	cell.Value = tv.Outputmap["time"]
	cell = row.AddCell()
	cell.Value = tv.Outputmap["触手TV"]
	cell = row.AddCell()
	cell.Value = tv.Outputmap["飞云TV"]
	cell = row.AddCell()
	cell.Value = tv.Outputmap["大神TV"]
	cell = row.AddCell()
	cell.Value = tv.Outputmap["企鹅电竞"]
	err = file.Save("data2.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println("触手TV:", tv.Outputmap["触手TV"])
	fmt.Println("飞云TV:", tv.Outputmap["飞云TV"])
	fmt.Println("大神TV:", tv.Outputmap["大神TV"])
	fmt.Println("企鹅电竞:", tv.Outputmap["企鹅电竞"])
}

func output2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("output recovered in f", r)
		}
	}()
	f, err := os.OpenFile("data.xls", os.O_APPEND, os.ModeAppend)
	if err != nil {
		if err != nil {
			f, err = os.Create("data.xls")
		}
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	if tv.First {
		w.Write([]string{"时间", "触手TV", "飞云TV", "大神TV", "企鹅电竞"})
		w.Flush()
		tv.First = false
		return
	}
	w.Write([]string{tv.Outputmap["time"], tv.Outputmap["触手TV"], tv.Outputmap["飞云TV"], tv.Outputmap["大神TV"], tv.Outputmap["企鹅电竞"]})
	w.Flush()
	fmt.Println("触手TV:", tv.Outputmap["触手TV"])
	fmt.Println("飞云TV:", tv.Outputmap["飞云TV"])
	fmt.Println("大神TV:", tv.Outputmap["大神TV"])
	fmt.Println("企鹅电竞:", tv.Outputmap["企鹅电竞"])
}
