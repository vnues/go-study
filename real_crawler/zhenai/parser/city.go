package parser

import ("regexp"

"../../engine"
"fmt"
)

const cityRe=`<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`


func ParseCity(contents []byte) engine.ParseResult{
	//fmt.Println(string(contents))
	re :=regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)
	result :=engine.ParseResult{}
	fmt.Println("启动城市详情页解析器")
	for _,m :=range matches {

		result.Items = append(result.Items, "User :"+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParseFunc: engine.NilParser,
		})
	}
		return result
}