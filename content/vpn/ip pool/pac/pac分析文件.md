### pac在win上的设置


	github.com/getlantern/pac


执行的是，使用自动配置脚本,localhost:1080/pac：

	func main() {
	pacUrl := `localhost:1080/pac`
	//https://github.com/getlantern/pac-cmd
	cmd := exec.Command(`C:\Users\mdshi\AppData\Roaming\byteexec\pac-cmd.exe`, "on", pacUrl)
	//设置代理
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	}


相当于：

	func main3() {
	helperFullPath := "pac-cmd"
	iconFullPath, _ := filepath.Abs("./icon.png")
	log.Debugf("Using icon at %v", iconFullPath)
	err := pac.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
	if err != nil {
		fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
		return
	}
	err = pac.On("localhost:12345/pac")
	if err != nil {
		fmt.Printf("Error set proxy: %s\n", err)
		return
	}
	fmt.Println("proxy set, Enter continue...")
	var i int
	fmt.Scanf("%d\n", &i)
	pac.Off("localhost:12345/pac")
	}


全代码：



	package main
	
	import (
		"fmt"
		"github.com/getlantern/golog"
		"github.com/getlantern/pac"
		"os/exec"
		"path/filepath"
	)
	
	var log = golog.LoggerFor("example")
	
	func main() {
		pacUrl := `localhost:1080/pac`
		//https://github.com/getlantern/pac-cmd
		cmd := exec.Command(`C:\Users\mdshi\AppData\Roaming\byteexec\pac-cmd.exe`, "on", pacUrl)
		//设置代理
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	
	}
	
	func main3() {
		helperFullPath := "pac-cmd"
		iconFullPath, _ := filepath.Abs("./icon.png")
		log.Debugf("Using icon at %v", iconFullPath)
		err := pac.EnsureHelperToolPresent(helperFullPath, "Input your password and save the world!", iconFullPath)
		if err != nil {
			fmt.Printf("Error EnsureHelperToolPresent: %s\n", err)
			return
		}
		err = pac.On("localhost:12345/pac")
		if err != nil {
			fmt.Printf("Error set proxy: %s\n", err)
			return
		}
		fmt.Println("proxy set, Enter continue...")
		var i int
		fmt.Scanf("%d\n", &i)
		pac.Off("localhost:12345/pac")
	}


参考文档:

[https://github.com/getlantern/pac-cmd](https://github.com/getlantern/pac-cmd)
[https://github.com/getlantern/byteexec](https://github.com/getlantern/byteexec)
[https://github.com/getlantern/pac](https://github.com/getlantern/pac)
[https://github.com/dawei101/tongsheClient.shadowsocks-go-ui](https://github.com/dawei101/tongsheClient.shadowsocks-go-ui)