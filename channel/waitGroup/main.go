package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, wg *sync.WaitGroup) {

	for n := range c {
		//并行打印
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}

}

type worker struct {
	in chan int
	//需要定义成引用类型 因为我们得去修改它得值
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i := 0; i < 10; i++ {

		workers[i].in <- 'a' + i

	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	//等待多人执行完
	//相当于time.Sleep 如果不执行完就给阻塞
	wg.Wait()

}

func main() {

	chanDemo()
}
