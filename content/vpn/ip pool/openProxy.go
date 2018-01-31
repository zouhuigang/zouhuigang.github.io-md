//go get golang.org/x/sys/windows/registry
/*
开启windows代理
https://github.com/jothep/maghat/blob/a7b3a8f8700ee2e9a142dd8467ef7da75d939699/maghat01.go
https://github.com/andrewgause/gostuff/blob/aee2e09bbb0a2a433bfe578a154367e46717b12d/main.go
https://studygolang.com/articles/8525
https://github.com/OuSatoru/djocore/blob/798ed4f9fb6d81181e6252af697b43fafa71cf2a/setIE/setie.go
https://github.com/anran800/registry/tree/f40da4ef4ac73c69a132f6062e6016b49035ef79
https://github.com/koofr/autoproxy/blob/f3cf5d14ef905fbbba90f1083ce09d2d0ef9b0b7/autoproxy_windows.go
https://pac.itzmx.com/
*/
package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

func main() {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	key.SetDWordValue("ProxyEnable", 0x00000000)

	//设置代理服务器ip和端口,ProxyServer ,183.133.66.72:8118
	/*

		// 写入32位整形值
		key.SetDWordValue("DWORD", 0xFFFFFFFF)

		// 写入64位整形值
		key.SetQWordValue("QDWORD", 0xFFFFFFFFFFFFFFFF)

		// 写入字符串
		key.SetStringValue("String", "hello")

		// 写入多行字符串
		key.SetStringsValue("Strings", []string{"hello", "world"})

		// 写入二进制
		key.SetBinaryValue("Binary", []byte{0x11, 0x22})
	*/

	defer key.Close()

	fmt.Printf("Windows proxy is closed.")

}

/*package main

import "golang.org/x/sys/windows/registry"
import "log"
import "fmt"

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatalf("opening registry: %v", err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
		log.Fatalf("ProxyServer key: %v", err)
	}
	fmt.Printf("Windows ProxyServer is %q\n", s)

	i, _, err := k.GetIntegerValue("ProxyEnable")
	if err != nil {
		log.Fatalf("ProxyEnable key: %v", err)
	}
	fmt.Printf("Windows ProxyEnable is %v\n", i)

	info, err := k.Stat()
	if err != nil {
		log.Fatalf("Getting Stats: %v", err)
	}
	fmt.Printf("Last modification time: %q\n", info.ModTime())

}*/
