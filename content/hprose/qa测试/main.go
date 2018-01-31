package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "1.xlsx" //yyyy-mm-dd hh:mm:ss
	arr := ReadExcel(excelFileName, 0, 5, 2)
	fmt.Printf("return data:%v\n", arr)

}

//这种方法，时间又问题
func ReadExcel2(filepath string, sheet int, startRow int, rowNum int) [][]string {
	if sheet < 0 {
		sheet = 0
	}
	if startRow < 0 {
		startRow = 0
	}
	if rowNum < 0 {
		rowNum = 10
	}
	endRow := startRow + rowNum

	//打开文件
	arr, err := xlsx.FileToSlice(filepath)
	if err != nil {
		cs := C.CString("open excel file error\n")
		C.myprint(cs)
	}
	arr1 := arr[0][startRow:endRow]
	return arr1
}

func ReadExcel(filepath string, sheet int, startRow int, rowNum int) [][]string {
	if sheet < 0 {
		sheet = 0
	}
	if startRow < 0 {
		startRow = 0
	}
	if rowNum < 0 {
		rowNum = 10
	}
	endRow := startRow + rowNum

	//打开文件
	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		cs := C.CString("open excel file error\n")
		C.myprint(cs)
	}
	var Data [][]string
	//循环 sheet
	for sk, sheet := range xlFile.Sheets {
		if sk > 0 {
			break
		}
		//定义slice 类型interface
		data := make([][]string, len(sheet.Rows))
		for k, row := range sheet.Rows {
			arr := make([]string, len(sheet.Rows[0].Cells))
			for s, cell := range row.Cells {
				str := cell.String()

				arr[s] = str
			}
			data[k] = arr
		}
		Data = data
	}

	/*arr, err := xlsx.FileToSlice(filepath)
	if err != nil {
		cs := C.CString("open excel file error\n")
		C.myprint(cs)
	}*/
	arr1 := Data[startRow:endRow]
	return arr1
}
