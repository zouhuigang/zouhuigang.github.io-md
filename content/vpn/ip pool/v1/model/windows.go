package model

import (
	"fmt"
	"syscall"
)

//打印彩色日志,SetConsoleTextAttribute函数是靠一个字节的低四来控制前景色，高四位来控制背景色
func ColorPrintln(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i)) //12 Red light

	fmt.Println(s)

	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7)) //White dark
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
