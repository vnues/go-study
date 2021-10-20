package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n2 int64 = 10
	fmt.Printf("n1的类型 %T n2占用的字节数是%d", n2, unsafe.Sizeof(n2))
}
