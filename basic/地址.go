package main

//引用必须使用不然会报错
import  "fmt"



// func main(){
// 	 a :=1
// 	 a=2
// 	 fmt.Println(a)
// }



func swap(a,b int){
	 a,b=b,a
	 a=6
}
//注意 前提是函数参数传值 但是 a :=1 a=2 还是会改变 因为这又不是参数传递 是赋值

func main(){
	 a,b:=4,5
	 swap(4,5)
	 fmt.Println(a,b)
}