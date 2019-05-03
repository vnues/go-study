package initt

import (
	"fmt"
)

var a int = b + c

var b int = f()
var c int = 1

func f() int {
	return c + 1
}

func Init() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
