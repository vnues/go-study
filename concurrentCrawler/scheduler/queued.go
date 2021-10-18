package scheduler

import (
	"concurrentCrawler/engine"
)

// QueuedScheduler TODO
type QueuedScheduler struct {
	// a int
	// b []int
	requestChan chan engine.Request // 这个channel用来放种子
	// workChan是chan类型 channel管道存放的类型是 chan engine.Request 所以有关workChan输入输出的操作，类型需要chan engine.Request
	workChan chan chan engine.Request // 这个channel用来存放包含种子的通道

}

// Submit TODO
func (s *QueuedScheduler) Submit(r engine.Request) {
	// 这里为什么不要初始化
	// s.workChan=make(chan  chan engine.Request)
	// 因为在run函数已经初始化了
	s.requestChan <- r
}

// ConfigureMasterWorkerChan TODO
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {}

// Run TODO
func (s *QueuedScheduler) Run() {
	/*
		fmt.Printf("----:%v",s.workChan)//nil
		s.workChan=make(chan  chan engine.Request)
		fmt.Printf("----:%v",s.workChan)//0xc000140000
		fmt.Printf("----:%v",s.a) //0
		fmt.Printf("----:%v",&s.a) //0xc0000ae480
		fmt.Printf("----:%v",s.b) //[]
	*/
	go func() {

		s.workChan = make(chan chan engine.Request)
		s.requestChan = make(chan engine.Request)
		// 创建两个队列用来存request和worker
		var requestQ []engine.Request // request队列
		var workerQ []chan engine.Request
		// 这两个队列的执行机制是这样的 request从队列拿出来然后交给worker队列拿出来的worker，workera需要这个种子才能执行  request->worker
		// request->worker是个channel操作,因为worker是channel channel的值操作就是输入输出
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 首选判断两个队列有没有
			if len(requestQ) > 0 && len(workerQ) > 0 {

				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			// 所有的channel操作都放在select select就是对channel操作的监听
			// 因为一个goroutine有多个channel 独立的channel所以用select来管理
			select {
			// 有新的request进来,放进去队列
			case request := <-s.requestChan:
				requestQ = append(requestQ, request)
				// 有新的worker进来，放进去队列
			case worker := <-s.workChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()

}

// 我们一个goroutine通常都是需要无限循环的

// WorkerReady TODO
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// 这个是创建一个可以接收种子的worker,需要放进去worker队列
	s.workChan <- w
}

// struct 的变量字段不能使用 := 来赋值以使用预定义的变量来避免解决 z.m=2
// 我的猜测 struct是的成员是channel类型是没有开辟新内存的

// worker队列就是in队列  用in队列映射worker队列

// 指针概念是全局的 你可以指针去找到对应到变量 不管这个变量在哪个作用域下

// struct里好像没有给channel类型开辟内存

// js 没指针的概念

// 指针 ->指向的内存

// slice声明的时候就分配了内存的

// Golang的引用类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性

// 引用类型就是通过指针去指向它这个变量内存

// 对 javascript确实没有指针对概念 指针很舒服 因为我可以不用考虑变量作用域问题

// 也就是说引用类型需要声明和开辟内存的
// 引用类型 声明引用类型 ->开辟内存然后分配指针给这变量 指针指向开辟的内存
// javascript是没有指针概念的-所以它也不需要我们可能自己手动分配内存
