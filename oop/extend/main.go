package main

import "fmt"

func main(){
	type Good struct {
		Name string
		Price int
	}
	type Brand struct {
		Name string
		Address string
	}

	type Tv struct {
		 *Good
		 *Brand
		 int //相当于int int名字和类型重名
	}

	tv:=Tv{
		&Good{
			Name:"电视机003",
			Price:199,
		},
		&Brand{
			Name:"海尔",
			Address:"广东",
		},
		//10,
		10,

	}
	// fmt.Println((*tv).Good) // 错误写法
	fmt.Println(*(tv.Good))
	fmt.Println(*(tv.Brand))
	fmt.Println(tv.int)

}

