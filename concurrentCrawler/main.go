package main

import (
	"concurrentCrawler/engine"
	"concurrentCrawler/scheduler"
	"concurrentCrawler/zhenai/parse"
	"fmt"
)

func main() {
	fmt.Println("开始启动项目")
	// 实例化一个QueueScheduler{}空结构体
	queueScheduler := scheduler.QueuedScheduler{}
	// 实例化一个ConcurrentEngine结构体
	e := engine.ConcurrentEngine{
		// 给它传QueuedScheduler 意思是QueuedScheduler有实现这个方法
		// 将QueuedScheduler{}空结构体指针传递过去
		// 使用指针方式传递，可以做到共享这个queueScheduler
		Scheduler:   &queueScheduler,
		WorkerCount: 100,
	}
	seed := engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.CityList,
	}
	// e.Run(seed1,seed2)
	// 为什么不&seed传递过去
	e.Run(seed)
}
