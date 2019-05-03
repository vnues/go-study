package main

import "fmt"

func main(){
	 var a int =20
	 
	 var ip *int
	 fmt.Println(ip)
     //fmt.Println(*ip)

	 ip = &a

	 fmt.Println(a)
	 fmt.Println(ip)
     fmt.Println(*ip)
}