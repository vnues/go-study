package engine

import "fmt"

// ConcurrentEngine TODO
// 定义管理并发版的engine
// engine 入口 ->它需要sheduler 和run
type ConcurrentEngine struct {
	// 这个就是接口的引用者
	Scheduler Scheduler
	// 定义开启worker的数量
	WorkerCount int
}

// Scheduler TODO
// 如果实现接口的方法的引用者是指针
// 那么我们用这个这个接口去当类型的时候，实例化接口变量也需要指针类型
type Scheduler interface {
	// Submit 接口定义的方法可以不用指明参数名
	// in<-r
	Submit(Request)
	// ConfigureMasterWorkerChan 这个接口的方法就是能让实现这个接口方法的人拿到in管道(in管道又放Requesr)
	ConfigureMasterWorkerChan(chan Request)

	Run()

	WorkerReady(chan Request)
}

// Run TODO
func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 定义两个管道
	// in :=make(chan Request) //Request类型的chan
	out := make(chan ParseResult)
	// 给实现接口方法的结构体拿到in
	// e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()
	// 开启多个createWorker--嗷嗷待哺要从in管道拿到种子
	for i := 0; i < e.WorkerCount; i++ {
		// createWorker需要从in管道拿到种子然后解析以后把结果放进去out管道
		// 传给它这个方法
		createWorker(e.Scheduler, out)
	}
	// 要把种子送机skeduler队列的方法--种子是通过in channel去拿的
	// channel:这段放后面是提前跟别人说你要在这里等着拿数据
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	// 记数
	itemCount := 0
	// 这个我们希望它也是无时无刻的去执行
	for {
		result := <-out
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
		// 打印返回来的parseResult
		for _, item := range result.Items {
			fmt.Printf("got item: %d %v\n", itemCount, item)
			itemCount++
		}

	}
}

// chan也是需要类型的指什么管道类型
// 管道的通信需要建立起码两个gorutine所以createWorker是个gorutine
// 开启一个协程
func createWorker(s Scheduler, out chan ParseResult) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
