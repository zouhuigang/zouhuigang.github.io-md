package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	. "github.com/bitly/go-simplejson"
	"github.com/satori/go.uuid"
	"github.com/zouhuigang/package/zcrypto"
	"github.com/zouhuigang/package/ztime"
)

/*
使用说明

第一次运行请先实例化对象（）

*/

type OperationInterface struct {

	//所有的接口操作都在这个结构体下
	client_id     string
	client_secret string
	access_token  string
	refresh_token string
}

type OurApplication struct {
	/*
		自有应用
		通过继承OperationInterface获得所有的操作方法，自身只有获取token的方法
	*/
	OperationInterface
}

type OpenApplication struct {
	/*
		开放应用
		通过继承OperationInterface获得所有的操作方法，自身只有获取token的方法
		code 需要手工获取
	*/
	OperationInterface
	code string
}

func (self OurApplication) AddPrinter(machine_code string, msign string) {
	url := "https://open-api.10ss.net/printer/addprinter"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&msign=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, msign)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) Sign(push_time int64) string {
	s := fmt.Sprintf("%s%d%s", self.client_id, push_time, self.client_secret)
	sign := zcrypto.Md5(s)
	return sign
}

func httpPost(data string, url string) (response string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Errorf("网络错误: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("数据解析错误：%s", err)
	}
	stats := string(body)
	return stats
}

func (self *OurApplication) GetToken() {
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	url := "https://open-api.10ss.net/oauth/oauth"
	fmtStr := fmt.Sprintf("cli状态码异常", err)
		return false
	}
	return true
}

//应用管理平台
//https://dev.10ss.net/admin/listinfo?id=1096322761
//接口文档-http://doc2.10ss.net/372519
func (self OperationInterface) Print(machine_code string, content string) {
	url := "https://open-api.10ss.net/print/index"
	uid := uuid.NewV4()
	timestamp := ztime.NowTimeStamp()
	origin_id := fmt.Sprintf("%d", uid)
	origin_id = zcrypto.Md5(origin_id)
	fmtStr := fmt.Sprintf("client_id=%s&access_token=%s&machine_code=%s&content=%s&origin_id=%s&sign=%s&id=%s&timestamp=%d",
		self.client_id, self.access_token, machine_code, content,
		origin_id, self.Sign(timestamp), uid, timestamp)
	fmt.Println(fmtStr)

	resp := httpPost(fmtStr, url)

	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) DelPrint(machine_code string) {
	url := "https://open-api.10ss.net/printer/deleteprinter"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) AddPrintMenu(machine_code string, content string) {
	url := "https://open-api.10ss.net/printmenu/addprintmenu"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&content=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, content)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) ShutdownRestart(machine_code string, response_type string) {
	url := "https://open-api.10ss.net/printer/shutdownrestart"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, response_type)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) SetSound(machine_code string, response_type string, voice string) {
	url := "https://open-api.10ss.net/printer/setsound"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s&voice=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, response_type, voice)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) PrintInfo(machine_code string) {
	url := "https://open-api.10ss.net/printer/printinfo"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) GetVersion(machine_code string) {
	url := "https://open-api.10ss.net/printer/getversion"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) CancelAll(machine_code string) {
	url := "https://open-api.10ss.net/printer/cancelall"
	uid := uuid.NewV4()
	timestamp := ztime.NowTimeStamp()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) CancelOne(machine_code string, order_id string) {
	url := "https://open-api.10ss.net/printer/cancelone"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&order_id=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, order_id)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) Seticon(machine_code string, img_url string) {
	url := "https://open-api.10ss.net/printer/seticon"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&img_url=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, img_url)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) DeleteIcon(machine_code string) {
	url := "https://open-api.10ss.net/printer/deleteicon"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) BtnPrint(machine_code string, response_type string) {
	url := "https://open-api.10ss.net/printer/btnprint"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&sign=%s&timestamp=%d&id=%s&access_token=%s&machine_code=%s&response_type=%s", self.client_id, self.Sign(timestamp), timestamp, uid, self.access_token, machine_code, response_type)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

func (self OperationInterface) GetOrder(machine_code string) {
	url := "https://open-api.10ss.net/printer/getorder"
	uid := uuid.NewV4()
	timestamp := time.Now().Unix()
	fmtStr := fmt.Sprintf("client_id=%s&access_token=%s&machine_code=%s&response_type=%s&sign=%s&id=%s&timestamp=%d",
		self.client_id, self.access_token, machine_code, "close", self.Sign(timestamp), uid, timestamp)
	fmt.Println(fmtStr)
	resp := httpPost(fmtStr, url)
	if self.CheckResponseStatus(resp) {
		fmt.Println(resp)
	} else {
		fmt.Println(resp)
	}
}

//订单完成
func finishOrd() {
	//cmd=oauth_finish&machine_code=1212&order_id=613&state=1&print_time=1426908344&origin_id=15966&push_time=1426908399&sign=1F19C52B0EE3FE0F36FEF7487795F9F7
}

func main2() {
	client_id := "1096322761"
	client_secret := "0b98b2a5341a5da3762cd20675bc9e95"

	//接单成功
	test := OurApplication{OperationInterface{client_id, client_secret, "9a9383d873e941cca815300e781c5126", "0e3824be65e3454e817bd4b113d166b7"}} //未获取到token时初始化空字符串，获取到之后再填进来

	test.CancelAll("4004545322")

}
func main() {

	/*
		初次运行请先调用getToken方法获取access_token及refresh_token，如无意外两个token会输出到命令行，复制两个token放入实例中使用
		未获取到token时初始化空字符串，获取到之后再填进来
	*/
	//应用id,应用密钥
	//test := OurApplication{OperationInterface{"1096322761", "0b98b2a5341a5da3762cd20675bc9e95", "", ""}} //未获取到token时初始化空字符串，获取到之后再填进来
	//open := OpenApplication{OperationInterface{"1096845322", "1c498e2c9214c0b712777b572ca2831f", "71ed2a877be6a820ddf86ddd0af9f528", "27d076a292ce8b0708972395850d7595"}, "d92b3a0eb2c1a33bb625e619e40be917"}
	//test.GetToken()
	client_id := "1096322761"
	client_secret := "0b98b2a5341a5da3762cd20675bc9e95"
	test := OurApplication{OperationInterface{client_id, client_secret, "9a9383d873e941cca815300e781c5126", "0e3824be65e3454e817bd4b113d166b7"}} //未获取到token时初始化空字符串，获取到之后再填进来
	test.AddPrinter("4004545322", "srkfnt5ytjr4")

	//学优教育辅导签到单打印===========
	title := `<FS><center>学优教育1v1辅导签到单</center></FS>`
	date := "\r\n辅导科目:语文\r\n辅导日期:2017-10-28\r\n辅导时段:08:00-10:00\r\n"
	name := "学生:何赵懿\r\n教师:李闻达 2017.10.28 14:21\r\n"
	qrcode := "<center> </center><QR>http://www.anooc.com/edu/teacher/sign?qrcode=29A9617D6D26B608</QR>注:此二维码有效期3天，过期作废!\r\n<center> </center>"
	sign := "学生签名:\r\n\r\n\r\n-----------------------------\r\n"
	bottomTips := "SCH143-29A9617D6D26B608\r\n*****请务必确认上课信息是否正确\r\n"
	content := zcrypto.Urlencode(title + date + name + qrcode + sign + bottomTips)
	test.Print("4004545322", content) // 打印接口
	// open.ShutdownRestart("400451758","restart")
}
