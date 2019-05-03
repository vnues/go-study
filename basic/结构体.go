package main

import "fmt"

type Vnues struct{
	 X int
	 Y int
}

// func main(){
// 	 fmt.Println(Vnues{4,5})
// }

func main(){
	 v := Vnues{1,4}
	 v.X = 4
	 v.Y=6
	 fmt.Println(v)
}