package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"github.com/zouhuigang/package/zsurname"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/*const (
	configFile = "config.txt"
)
*/
var (
	p   zsurname.ConfigDat
	err error
)

func loadConfig() {
	p, err = zsurname.New()
	if err != nil {
		log.Println("config file load fail ..\n")
	}
	log.Println("config reload success .......\n")

}

func listen() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1) //SIGINT SIGUSR1
	for {
		<-c
		loadConfig()
	}
}

//返回结构
type Result struct {
	Success int
	Msg     string
}

func findsurname(realname string) *Result {
	success, msg := p.FindSurname(realname)
	if success {
		return &Result{1, msg}
	}
	return &Result{0, "[" + realname + "]" + msg}
}

func main() {
	pid := os.Getpid()
	log.Printf("PID: %d\n", pid)
	go listen()
	loadConfig() //加载配置
	//rpc server
	service := rpc.NewHTTPService()
	service.AddFunction("findsurname", findsurname, rpc.Options{})
	http.ListenAndServe(":3335", service)

	/*for {
	}*/
}
