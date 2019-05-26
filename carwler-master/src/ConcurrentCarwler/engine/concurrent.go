package engine

import "log"

//声明ConcurrentEngine对象
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//Scheduler接口 需要有人去实现这个接口
type Scheduler interface {
	Submit(r Request)
	ConfigureMasterWorkerChan(c chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	//初始化调度器chan
	e.Scheduler.ConfigureMasterWorkerChan(in)

	//创建多个worker
	for i := 0; i <= e.WorkerCount; i++ {
		createWorker(in, out)
	}
	//把所有的请求提交到chan中
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	//无限循环 一直在拿到爬取的信息和将种子放进in管道（队列）
	//就是不会停一直在走
	//等到in out双方完成自己的动作 就走下次循环？
	for {
		//out管道有值 就走 不然卡住不动，不会继续循环知道有值
		//与无限for循环的妙用
		result := <-out
		//什么时候会成功执行这一步
		//就是in能把种子发给其他worker接收
		//因为in一阻塞 out也不会执行
		//所以得等待
		//原因是他们都处于同一个gorutine函数
		//简单来说两个管道有影响到了
		//result什么时候能收啊就是			e.Scheduler.Submit(request)全部能送走
		for _, item := range result.Items {
			log.Printf("Got item: &%d,%v", itemCount, item)
			itemCount++
		}
		/*
			Requests []Request
		*/
		//你怎么保证就是in管道不会有缓存（积压）

		//in管道只放一个request而不是request数组这一点得知道
		for _, request := range result.Requests {
			/*
			func (s *SimpleScheduler) Submit(r engine.Request) {
				//go func() {
					//in引用
					//in管道接收
					s.workerChan <- r
				//}()
			}
			*/
			//没有开启gorutine就会造成死锁
			e.Scheduler.Submit(request)
		}
	}
}

/**
每个worker开启一个协程 等待chan的数据返回
*/
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		//一直执行worker
		for {
			//in管道发送
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}

//问题一in out 管道都是要保持为没有缓存的 就是你有接收就要发送
//out parseResult接收到就发送
/*
	Requests []Request
*/
//你怎么保证就是in管道不会有缓存（积压）

//in管道只放一个request而不是request数组这一点得知道