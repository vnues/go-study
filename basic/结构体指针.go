package main

import "fmt"

type Vnues struct{
	X int
	Y int
}


func main(){
	  v := Vnues{5,6}
	  p :=&v
	  p.X=1e9
	  fmt.Println(v)
}