/*
windows下配置pac代理
//在windows下->控制面板->Internet选项->连接->局域网设置->使用自动配置脚本->地址填入http://127.0.0.1:8888/pac/pac.txt

*/
package model

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"golang.org/x/sys/windows/registry"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unsafe"
)

const (
	//win+r ->regedit->
	WinPathInternetSettings = `Software\Microsoft\Windows\CurrentVersion\Internet Settings`
	KEYNAME                 = "AutoConfigURL"
)

func safeParseUrl(in string) (*url.URL, error) {
	if !strings.HasPrefix(in, "http") {
		in = fmt.Sprintf("http://%s", in)
	}
	return url.Parse(in)
}

func EditReg(KEYVALUE string) {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, WinPathInternetSettings, registry.ALL_ACCESS)

	if err != nil {
		log.Fatal(err)
	}

	defer key.Close()
	err = key.SetStringsValue(KEYNAME, []string{KEYVALUE})
	if err != nil {
		log.Fatal(err)
	}
	//refresh
	refreshReg()
	msgChan := make(chan os.Signal, 1)
	signal.Notify(msgChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGQUIT)
	<-msgChan
	func() {
		fmt.Println("Application end")
		err = key.DeleteValue(KEYNAME)
	}()
}

//刷新注册表
func refreshReg() {

	firstP := int64(0)
	internetOptionProxy := int64(39)
	internetOptionProxy1 := int64(37)
	internetInfo := int64(0)
	sizeOfInfo := int64(0)
	wininet, _ := syscall.LoadLibrary("wininet.dll")
	defer syscall.FreeLibrary(wininet)
	InternetSetOption, _ := syscall.GetProcAddress(syscall.Handle(wininet), "InternetSetOptionA")
	r, _, _ := syscall.Syscall6(uintptr(InternetSetOption), 4,
		uintptr(firstP),
		uintptr(internetOptionProxy),
		uintptr(unsafe.Pointer(&internetInfo)),
		uintptr(sizeOfInfo), 0, 0)
	if r != 0 {
		fmt.Println("Zproxy internetOptionProxy has been successful in modifying the registry..")
	}
	r1, _, _ := syscall.Syscall6(uintptr(InternetSetOption), 4,
		uintptr(firstP),
		uintptr(internetOptionProxy1),
		uintptr(unsafe.Pointer(&internetInfo)),
		uintptr(sizeOfInfo), 0, 0)
	if r1 != 0 {
		fmt.Println("Zproxy internetOptionProxy1 has been successful in modifying the registry...")
	}

}

//设置自动配置pac文件地址
func SetPac() {
	key, err := registry.OpenKey(registry.CURRENT_USER, WinPathInternetSettings, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	//key.SetDWordValue("ProxyEnable", 0x00000000)
	// 写入字符串
	key.SetStringValue("AutoConfigURL", "http://127.0.0.1:9999/pac/pac.txt")

	//设置代理服务器ip和端口,ProxyServer ,183.133.66.72:8118
	defer key.Close()

	refreshReg()
}

//读取pac文件
func ProxyFromAutoConfig(req *http.Request) (proxyUrl *url.URL, err error) {

	// check if AutoConfigURL is definded in registry, otherwise no point in spending time to check it
	k, err := registry.OpenKey(registry.CURRENT_USER, WinPathInternetSettings, registry.READ)
	if err != nil {
		return
	}
	defer k.Close()

	autoConfigURL, _, err := k.GetStringValue("AutoConfigURL")
	if err != nil {
		// no auto proxy url
		return nil, nil
	}

	res, err := http.Get(autoConfigURL)
	if err != nil {
		return nil, err
	}

	vm := otto.New()
	vm.Run(res.Body)
	vm.Set("url", req.URL)
	vm.Set("host", req.URL.Host)
	vm.Run("proxy = FindProxyForURL(url, host);")
	val, err := vm.Get("proxy")
	if err != nil {
		return nil, err
	}

	proxyStr, err := val.ToString()
	if err != nil {
		return nil, err
	}
	// parse it
	// format reference: https://en.wikipedia.org/wiki/Proxy_auto-config
	for _, proxyMaybe := range strings.Split(proxyStr, ";") {
		proxyMaybe = strings.ToLower(proxyMaybe)
		if strings.Contains(proxyMaybe, "direct") {
			return nil, nil
		}

		proxyMaybe = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(proxyMaybe), "proxy "))

		urlMaybe, err := safeParseUrl(proxyMaybe)
		if err == nil {
			return urlMaybe, err
		}
	}

	return
}
