package main

import (
	"./complex"
	initt "./init"
)

//别名
//使用的时候小心  如果上面的模块有使用别名 那么底下好像也要用 不是 init是特殊字段
func main() {
	// m.Println("function")

	// h.Hello()
	// h.World()
	//m.Println(h.Name)
	initt.Init()
	complex.Img()
	//m.Println(h.age)
}

//一个目录下只能放一个main包
