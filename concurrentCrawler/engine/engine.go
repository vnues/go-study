// Package engine TODO
package engine

import (
	"concurrentCrawler/fetch"
	"fmt"
)

// Run TODO
func Run(seeds ...Request) {
	fmt.Println("启动种子引擎")
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		ParseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			fmt.Printf("爬取所得到到信息：%v\n", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetch.Fetcher(r.Url)
	fmt.Printf("我正在爬取这个地址Fetching %s\n", r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	ParseResult := r.ParseFunc(body)
	return ParseResult, nil
}

// go append方式是slice方法 有两种形式追加和拼接(就是打散切片)

// engine这快是把爬取的数据打印出来并且又把种子放进去队列执行

/*
‘…’ 其实是go的一种语法糖。
它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
第二个用法是slice可以被打散进行传递
*/
/*
func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
}
*/
