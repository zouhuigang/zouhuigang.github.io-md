package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/zouhuigang/package/zphone/read"
	"net/http"
)

var zphoneRead read.PhoneST

func init() {
	zphoneRead = read.LoadDatFile(`phone.dat`)
}

type PhoneRecord struct {
	ResultCode int    //1解析成功,0解析失败
	Msg        string //错误原因
	PhoneNum   string
	Province   string
	City       string
	ZipCode    string
	AreaZone   string
	CardType   string
}

func findPhone(mobile string) *PhoneRecord {
	pr, err := zphoneRead.Find(mobile)
	if err != nil {
		fmt.Printf("%s,%s\n", mobile, err)
		return &PhoneRecord{0, err.Error(), mobile, "", "", "", "", ""}

	}

	fmt.Printf("%s,%s,%s,%s,%s,%s\n", mobile, pr.Province, pr.City, pr.ZipCode, pr.AreaZone, pr.CardType)

	return &PhoneRecord{1, "解析成功", mobile, pr.Province, pr.City, pr.ZipCode, pr.AreaZone, pr.CardType}
}

func main() {
	service := rpc.NewHTTPService()
	//service.AddFunction("queue", queue, rpc.Options{})
	service.AddFunction("findPhone", findPhone, rpc.Options{})
	http.ListenAndServe(":3333", service)
}
