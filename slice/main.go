package main

import (
	"fmt"
)

func PrintArray(array []int) {
	for _, v := range array {
		fmt.Println(v)
	}
}

//	切片是对数组底层对view
func main() {
	var a = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := a[3:6]
	PrintArray(s1) //[3,4,5]
	s2 := s1[3:4]  //[6]
	PrintArray(s2)
}
