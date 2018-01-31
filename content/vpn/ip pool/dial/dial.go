//https://www.cnblogs.com/yjf512/archive/2012/06/16/2552296.html
package main

import (
	"net"

	"fmt"
)

func main() {

	currency := 20 //并发数,记住，一个连接数是打开一个端口号，window和linux的端口号都是有限制的

	count := 10 //每条连接发送多少次连接

	for i := 0; i < currency; i++ {

		go func() {

			for j := 0; j < count; j++ {

				sendMessage()

			}

		}()

	}

	select {}

}

func sendMessage() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {

		panic("error")

	}

	header := "GET / HTTP/1.0\r\n\r\n"

	fmt.Fprintf(conn, header)

}
