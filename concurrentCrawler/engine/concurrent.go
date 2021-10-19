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
	// in :=make(chan Request)
	// 输出结果类型的channel
	// out是共用一个的
	out := make(chan ParseResult)
	// 给实现接口方法的结构体拿到in
	// e.Scheduler.ConfigureMasterWorkerChan(in)
	// 调度器初始化工作
	e.Scheduler.Run()
	// 开启多个createWorker--嗷嗷待哺要从in管道拿到种子
	// WorkerCount worker的数量
	for i := 0; i < e.WorkerCount; i++ {
		// createWorker需要从in管道拿到种子然后解析以后把结果放进去out管道
		// 传给它这个方法
		createWorker(e.Scheduler, out)
	}
	// 要把种子送去scheduler队列的方法--种子是通过in channel去拿的
	for _, r := range seeds {
		// 将request放入request channel
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
			// TODO 这里可以进行sql或者数据存储操作
			fmt.Printf("got item: %d %v\n", itemCount, item)
			itemCount++
		}
	}
}

// 函数参数规范
// 传入变量和返回变量以小写字母开头。
// 参数数量均不能超过5个。
// 尽量用值传递，非指针传递。
// 传入参数是 map，slice，chan，interface 不要传递指针。

// 管道的通信需要建立起码两个goroutine所以createWorker是个goroutine
// 开启一个协程
func createWorker(s Scheduler, out chan ParseResult) {
	// 缓存seed的channel
	// 每调用一个createWorker就创建一个in channel
	// 一个worker对应一个in channel
	in := make(chan Request)
	go func() {
		// 一个worker消耗一个in channel (存放多个request)
		// 再也不是多个worker去抢一个in channel
		for {
			// 一个worker准备好了
			// in - 存放request的channel
			// workChan是 worker channel
			// WorkerReady: s.workChan <- in
			// channel也对应一个make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil
			// 派生类型的函数参数传递是拷贝了一份地址 原来是这样
			// 这里就是的动作：其实就是让request传过来
			// 我worker准备好了,request可以传过来
			// TODO 但是worker获取到的request也是无序的
			s.WorkerReady(in)
			// request就是seed
			request := <-in
			// worker的工作包括fetch+parse
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			// 将结果放进去out channel
			out <- parseResult
		}
	}()
}
