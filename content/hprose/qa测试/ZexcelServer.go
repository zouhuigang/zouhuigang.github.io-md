package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/tealeg/xlsx"
	"github.com/zouhuigang/package/zcrypto"
	"net/http"
	"runtime"
	"runtime/debug"
)

const (
	APITOKEN = `73b0f08a945bb480d0096c70142bbe84` //验证通过API
)

func hello(name string) string {
	return "Hello " + name + "!"
}

//登录获取token
func login(user, passwd string) (int, string) {
	passwdMd5 := zcrypto.Md5UP(passwd)
	passConfirm := zcrypto.Md5UP(`zex123456`)
	fmt.Printf("login user : %v-%v\n", user, passwd)
	if user == "zex01" && passwdMd5 == passConfirm {
		return 1, APITOKEN
	}
	msg := fmt.Sprintf("用户[%s-%s]登录失败，无法获取接口秘钥!", user, passwd)
	return 0, msg
}

//核对登录
func checkLogin(apitoken string) (bool, string) {
	if apitoken == APITOKEN {
		return true, "登录api成功"
	}

	return false, "apitoken验证失败"
}

func readExcel(filepath string, sheetNum int, startRow int, apitoken string) [][]string {
	var arr [][]string

	//核对登录接口api
	is_success, msg := checkLogin(apitoken)
	if !is_success {
		fmt.Printf("error:%s\n", msg)
		panic("接口认证失败") //抛出错误，php可通过try catch捕获信息
		//return arr //这句不会执行
	}

	defer gc()
	if sheetNum < 0 {
		sheetNum = 0
	}
	if startRow < 0 {
		startRow = 0
	}

	//arr, err := xlsx.FileToSlice(filepath)//时间又问题
	//if err != nil {
	//fmt.Println()
	//}
	//打开文件
	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		fmt.Printf("open excel file error:%v\n", err)
	}

	//循环 sheet
	for sk, sheet := range xlFile.Sheets {
		if sk != sheetNum { //读取第几个sheet
			continue
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
		arr = data
	}

	arr1 := arr[startRow:]

	st := fmt.Sprintf("read excel success:%s-%d-%d", filepath, sheetNum, startRow)
	fmt.Println(st)
	return arr1
}

//释放内存
func gc() {
	debug.FreeOSMemory()
	runtime.GC()
}

func main() {
	service := rpc.NewHTTPService()

	service.AddFunction("login", login)
	service.AddFunction("hello", hello, rpc.Options{})
	service.AddFunction("readExcel", readExcel, rpc.Options{})
	http.ListenAndServe(":9999", service)
}
