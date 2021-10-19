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
	go func() {
		s.workChan = make(chan chan engine.Request)
		s.requestChan = make(chan engine.Request)
		// 创建两个队列用来存request和worker
		// 并且队列
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
				// worker和request都准备好了 就传过去 可以干活了
				// 其实最后是转化为 in <- activeRequest
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

// WorkerReady TODO
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// 这个是创建一个可以接收种子的worker,需要放进去worker队列
	s.workChan <- w
}
