package engine

import (
	"../fetcher"
	"fmt"
	"log"
)

//这个engine下的方法类型 不用加engine。

//seeds是种子
//参数多个类型一样的
//变长参数，参数数量不确定 类型都是request
//你可能会想到为什么不传数组
func Run(seeds ...Request){
	fmt.Println("启动engine")
     var requests []Request
     for _,r :=range seeds{
     	fmt.Println(r)
     	requests = append(requests,r)
	 }
     for len(requests)>0{
     	r :=requests[0]
     	requests =requests[1:]

     	body,err:=fetcher.Fetch(r.Url)
     	fmt.Println(body)
     	if err !=nil{
     		log.Printf("Fetcher:error"+"fetching url %s :%v",r.Url,err)
     		//跳过这次循环去执行下一步循环
     		continue
		}
		 ParseResult :=r.ParserFunc(body)
     	fmt.Println(ParseResult) //[] []}
     	requests = append(requests,ParseResult.Requests...)
     	for _,item :=range  ParseResult.Items{

			log.Printf("Got item %v",item)
		}
	 }
}