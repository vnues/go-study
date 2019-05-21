package parser

import (
	"../../engine"
	"fmt"
	"regexp"
)

const cityListRe =`<a target="_blank" href="(https://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<+])</a>`
//这个解析以后返回ParseRequest->包含request(请求)
func ParseCityList(contents []byte) engine.ParseResult{
	//* 匹配前一个字符0或无限次
	fmt.Println("启动解析器")
	fmt.Println(string(contents))
	  re :=regexp.MustCompile(cityListRe)
	  fmt.Println(re)

	matches := re.FindAllSubmatch(contents,-1)
	   fmt.Println(matches)//[]
	result :=engine.ParseResult{}
	  for _,m :=range matches{
	  	//要生成新的request
	  	//unc append(slice []Type, elems ...Type) []Type
        //没有匹配上 这个正则
	  	result.Items = append(result.Items,m[2])
	  	result.Requests = append(result.Requests,engine.Request{
	  		Url:string(m[1]),
	  		ParserFunc :engine.NilParser,
		})
	  }
	  return result
}