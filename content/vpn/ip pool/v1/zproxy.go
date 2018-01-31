/*
Z代理,windows下使用
爬取一些开放的代理地址，实现绕过屏蔽，翻墙的目的。
作者:邹慧刚 zouhuigang888@gmail.com
https://github.com/denghongcai/yaproxy.git
https://segmentfault.com/q/1010000000150166/
https://github.com/elazarl/goproxy/issues/192
https://www.cnblogs.com/zhangqingping/p/4344278.html
*/

package main

import (
	"./model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StaticServer(w http.ResponseWriter, req *http.Request) {
	staticHandler := http.FileServer(http.Dir("./pac/"))
	staticHandler.ServeHTTP(w, req)
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}
	if path == "" {
		path = "index.html"
	}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(w, "404")
		return
	}
	fmt.Fprintf(w, "%s\n", contents)
}

func closeMain(s1 os.Signal) {
	model.EditReg(``)
	fmt.Println("zproxy has been closed....\nexit signal:", s1)
}

func main() {
	/*fakeReq, _ := http.NewRequest("GET", "http://google.com", nil)
	p, err := ProxyFromAutoConfig(fakeReq)
	fmt.Printf("ProxyFromAutoConfig: %s %vn", p, err)*/

	model.ColorPrintln(".\n..\n...\n......\n.........\n......... Z代理开始运行，需要翻墙时，请勿关闭,占用本地端口:9999,9797\n.........\n......... 作者:邹慧刚 官网地址:www.anooc.com\n.........\n......\n...\n..\n.\n", 10)

	//修改注册表
	model.SetPac()

	//开启pac文件服务器
	http.HandleFunc("/", Handler)
	http.HandleFunc("/pac", StaticServer)
	s := &http.Server{
		Addr: ":9999",
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	//开启代理服务器,http://ip.chinaz.com/getip.aspx
	ip := model.GetRandomIp() //"http://61.155.164.111:3128"
	proxy, _ := model.ProxyHTTP(ip)
	//http.TimeoutHandler(1,
	//	1000*time.Millisecond, `dd`)
	zproxy := &http.Server{
		Addr:           ":9797",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        proxy,
	}
	go func() {
		log.Fatal(zproxy.ListenAndServe())
		//log.Fatal(http.ListenAndServe(":9797", proxy)) //proxy
	}()

	//捕获ctrl+c
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("zproxy is running....")
	s1 := <-c
	closeMain(s1)
}
