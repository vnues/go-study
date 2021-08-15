package main

import "fmt"

func main(){
   s:= "hello 中国"
   println(len(s)) // 12
   for _,b:=range []byte(s){
   	fmt.Printf("%X ",b) // 68 65 6C 6C 6F 20 E4 B8 AD E5 9B BD
   }
   // rune
   for _,ch :=range []rune(s){
   	fmt.Printf("%c ",ch) // h e l l o   中 国
   }
}
