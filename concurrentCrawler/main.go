package main

import (
	"concurrentCrawler/engine"
	"concurrentCrawler/scheduler"
	"concurrentCrawler/zhenai/parse"
	"fmt"
)

func main() {
	fmt.Println("开始启动项目")
	e := engine.ConcurrentEngine{
		// 给它传QueuedScheduler 意思是QueuedScheduler有实现这个方法
		// 我感觉跟实例化对象一样
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.CityList,
	})
}
