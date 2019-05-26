package scheduler

import (
	"ConcurrentCarwler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		//in引用
		//in管道接收
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	//channel是引用传递
	s.workerChan = c
}
