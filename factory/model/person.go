package model

import "fmt"

type  student struct {
	Name string
	age int
}

// 工厂函数
func NewStudent(n string,a int) *student{
	 return &student{
	 	Name:n,
	 	age:a,
	 }
}
// getter方法获取get


func AgePrint(){

	 var s = student{
	 	Name:"vnues",
	 	age:22,
	 }
	 fmt.Println(s.age)
}