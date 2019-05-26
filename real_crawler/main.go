package main

import (
	"./engine"
	"./zhenai/parser"
	"fmt"
)
func main(){
 fmt.Println("33")
  engine.Run(engine.Request{
  	Url:"http://www.zhenai.com/zhenghun",
  	ParserFunc:parser.ParseCityList,
  })
}
