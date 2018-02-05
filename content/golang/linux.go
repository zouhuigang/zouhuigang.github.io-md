package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

const (
	configFile = "config.txt"
)

func loadConfig() {
	bts, err := ioutil.ReadFile(configFile)
	if err != nil {
		// handle config file error
		fmt.Println(err.Error())
	} else {
		// do something with configurations
		fmt.Printf(string(bts))
	}
}

func listen() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	for {
		<-c
		loadConfig()
	}
}

func main() {
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)
	go listen()
	loadConfig()
	for {
	}
}
