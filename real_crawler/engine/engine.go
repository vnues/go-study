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
//seeds Request类型有多个 打印seeds是切片类型
func Run(seeds ...Request){
	fmt.Println("启动engine")
	fmt.Println(seeds)
	//拷贝这个对象 不能直接在里面操作 切片是引用传递
     var requests []Request
     for _,r :=range seeds{
     	requests = append(requests,r)
	 }
     for len(requests)>0{
     	r :=requests[0]
     	requests =requests[1:]
     	body,err:=fetcher.Fetcher(r.Url)
     	//fmt.Println(string(body))
     	fmt.Printf("Fetching %s \n",r.Url)
     	if err !=nil{
     		log.Printf("Fetcher:error"+"fetching url %s :%v",r.Url,err)
     		//跳过这次循环去执行下一步循环
     		continue
		}
     	 //返回的是ParseResult
		 ParseResult :=r.ParserFunc(body)
		 //把种子送进消息队列
     	requests = append(requests,ParseResult.Requests...)
     	for _,item :=range  ParseResult.Items{
			fmt.Printf("Got item %v\n",item)
		}
     	//for _,url :=range  ParseResult.Requests{
		//	fmt.Printf("Got url %v\n",url.Url)
		//}
	 }
}