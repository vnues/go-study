package main

import "fmt"

func main(){
	 //p :=  [3] int{2,3,4,5,6,67}
	//  p :=  [...] int{2,3,4,5,6,67}
	p :=  [] int{2,3,4,5,6,67} // 未指定容器大小的数组
	 fmt.Println("p===",p)
	 for i :=0;i<len(p);i++{
		 fmt.Printf("p[%d]==%d\n",i,p[i])
	 }
}