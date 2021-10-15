package main

import (
	"./model"
	"fmt"
)


func main(){
	// Golang 在创建结构体实例(variable)时，可以直接指定字段的值
	// 这样可以不用指定这个变量是什么类型 因为自动赋值就确定了 正常声明是var variable 变量类型
	var student =model.NewStudent("tom",26)
	fmt.Println(*student) // {tom 26}
	fmt.Println(student.Name) // tom
	// fmt.Println(student.age) // student.age undefined (cannot refer to unexported field or method age)
	model.AgePrint()
}


// 这里结构体里小写字段有疑问 在model包中我已经return回来 为什么还不能访问不到 我return回来这个结构体或者结构体指针赋值给这个变量
// 这个变量指向的结构体按道理在这个包应该全部都能访问得到

// 类似函数的private但是go中private属性本包是可以访问得到的 这就是跟java的区别

// 用工厂模式实现一个构造函数