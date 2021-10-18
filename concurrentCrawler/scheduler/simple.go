package scheduler

import (
	"concurrentCrawler/engine"
)

// 实现interface接口

// SimpleScheduler TODO
type SimpleScheduler struct {
	// 其实就是in管道
	workChan chan engine.Request
	// 也就是当需要改变SimpleScheduler的属性的时候就需要指针传递
}

// Submit TODO
// 要求引用者是指针类型并且符合SimpleScheduler类型
// 为什么传过去要指针类型 因为我们的workChan需要被操作 我们不希望拷贝的workChan来操作
func (s *SimpleScheduler) Submit(r engine.Request) {
	// 这里就是将Request种子放进in管道
	// 但是你有没有看到in管道变量是engine下的 不是这个包的
	// 所以我们得把in给拿过来
	go func() {
		s.workChan <- r
	}()
}

// ConfigureMasterWorkerChan TODO
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}
