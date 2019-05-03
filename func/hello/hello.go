package hello

//属于hello这个包

import (
	"fmt"
)

//大写开头的属性代表外部可以访问  但是小写代表是私有的
const Name string = "vnues"

const age int = 18

func Hello() {
	fmt.Println("hello")
}
