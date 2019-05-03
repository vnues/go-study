package main

import (
	"fmt"
)

//数组 数量写在类型的前面
//循环的时候 range关键字可以拿到下标和值
//如果我们不想要下标可以用_表示 可通过_ 省略变量
func main() {
	//int数组初始值为0 布尔值数组默认是false
	var a [3]int
	var b [3]bool
	//声明的时候赋值
	var c [3]int = [3]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for i, v := range b {
		fmt.Printf("%d %v\n", i, v)
	}
	for i, v := range c {
		fmt.Printf("%d %v\n", i, v)
	}
	//:= 赋值方式需要给初始值
	d := [3]int{3, 4, 5}
	for i, v := range d {
		fmt.Printf("%d %v\n", i, v)
	}
	//任意数组长度[...]让编译器来帮我们数多少
	//数组是值类型
	//[10]int [20]int是不同类型
	//调用fuc f(arr[10]int)会拷贝数组
	//以数组指针做参数传递过去
	e := [5]int{1, 2, 3, 4, 5}
	fmt.Println("-----------分割线---------------")
	//原来是这样表示
	//go语言的指针很灵活 数组都不用去转化
	fmt.Println((&e)[0])
	PrintArray(&e)
	//这样子很麻烦又要知道数组长度又要使用指针
	//go语言一般不会使用数组而是使用切片(slice)

}
func PrintArray(array *[5]int) {
	for i, v := range array {
		fmt.Printf("%d %v\n", i, v)
	}
}
