package main

import "fmt"

//函数作为返回参数的写法
func adder() func(int) int{
	  sum :=0

	  return func(v int) int{
	  	  sum=v+sum
	  	  return sum
	  }

}


func main(){
	a := adder()
     for  i:=0;i<10;i++{
     	 fmt.Printf("0+1+......+%d=%d\n",i,a(i))
	}

}

//闭包return一个函数

//只要是类型 就能对应的接口

//定义接收者的方法

//接收者方法也有一些限制，这也是它和扩展方法之间的区别。接收者方法的接受者类型，必须和接收者方法定义在同一个包中。所以很多非自定义的类型，以及基本类型都不能当做接收者的类型。当然也可以投机取巧，在自己的包中重新
//为这些类型取个名字即可。

//把基本类型重新定义一下，就可以当做接收者类型了
//type MyString string
//
//func (str MyString) hello() {
//	fmt.Println("hello" + str)
//}

//go是没有类型这个概念的
//
//实现接口
//在Golang中，其实并没有“实现接口”这一说法。在Golang中接口是隐式实现的，也就是说我们不需要implements这些关键字。
//只要一个类型的接收者方法和接口中定义的方法一致，Golang就认为这个类型实现了该接口。下面是一个简单的例子。、



//函数
//-闭包
//-函数实现接口