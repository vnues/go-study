package parse

import (
	"concurrentCrawler/engine"
	"fmt"
	"regexp"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th>`

// City TODO
func City(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	fmt.Println("启动城市详情页解析器")
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User :"+name)

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return Profile(c, name)
			},
		})
	}
	return result
}
