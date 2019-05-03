package main

import (
	 "fmt"
	 "math"
	 "reflect" 
)

func main(){
	 var x,y int =3,4
	 var f float64= math.Sqrt(float64(x*x+y*y))
	 fmt.Println(reflect.TypeOf(f))
	 fmt.Println(f)
	 var z int = int(f)
	 fmt.Println(reflect.TypeOf(z))
	 fmt.Println(z)

}