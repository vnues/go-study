// package main

// import "fmt"

// func main(){
// 	 //s := [] int{1,2,34,5}
// 	 s := make([]int,4,5)
// 	 fmt.Println(s)
// 	 // s[4]=4 报错
// 	 fmt.Println(s)
// }

package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%d\n",len(x),cap(x),x)
}