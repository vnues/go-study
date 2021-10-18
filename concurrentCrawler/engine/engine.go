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

// 怎么传参不用指针类型
func worker(r Request) (ParseResult, error) {
	body, err := fetch.Fetcher(r.Url)
	fmt.Printf("我正在爬取这个地址Fetching %s\n", r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	ParseResult := r.ParseFunc(body)
	return ParseResult, nil
}
