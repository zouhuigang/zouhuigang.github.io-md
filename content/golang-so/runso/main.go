package main

/*
#include "libdoubler.h"
#cgo LDFLAGS: -ldoubler
*/
import "C"

import (
	"fmt"
)

func main() {

	str := C.DoubleIt(21)
	fmt.Println(str)
}
