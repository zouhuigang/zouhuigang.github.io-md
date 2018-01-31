package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/zouhuigang/package/zip"
	"net/http"
)

type IpRecord struct {
	ResultCode int    //1解析成功,0解析失败
	Msg        string //错误原因
	Ip         string
	Continent  string //洲
	Country    string //国家
	Province   string //省份
	City       string //城市
	County     string //区县
	Operator   string //运营商
	Zoning     string //区划，区号
	Ecountry   string //国家英文
	Ecode      string //国家简码
	Long       string //经度
	Lat        string //纬度
}

func findIp(ip string) *IpRecord {
	zipRead, _ := zip.New()
	pr, err := zipRead.FindIp(ip)
	if err != nil {
		fmt.Printf("%s,%s\n", ip, err)
		return &IpRecord{0, err.Error(), ip, "", "", "", "", "", "", "", "", "", "", ""}

	}

	fmt.Printf("%s,%s,%s\n", ip, pr.Continent, pr.Country)

	return &IpRecord{1, "解析成功", ip, pr.Continent, pr.Country, pr.Province, pr.City, pr.County, pr.Operator, pr.Zoning, pr.Ecountry, pr.Ecode, pr.Long, pr.Lat}
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("findIp", findIp, rpc.Options{})
	http.ListenAndServe(":3334", service)
}
