package main

import "fmt"

func getVal(num1 int,num2 int)(int,int){
	 sum :=num1+num2
	 sub :=num2-num1
	 return sum,sub
}

func main(){

	sum,sub:=getVal(30,30)
	fmt.Println("sum=",sum,"sub=",sub)
	sum2,_:=getVal(10,30)
	fmt.Println("sum=",sum2)

	// 定义变量 声明变量
	var i int

	// 给i赋值
	i=10

	// 使用变量
	fmt.Println("i=",i)
}
