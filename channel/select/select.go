package main

import "fmt"

/*
如果存在多个 channel 的时候，通过 select 可以监听 channel 上的数据流动。

select 默认是阻塞的，只有当监听的 channel 中有发送或接收可以进行时才会运行，当多个 channel 都准备好的时候，select是随机的选择一个执行的。
*/
//定义数列 --->channel的写法
//对于读case语句，如果当前Channel有数据可读，则执行。
// 对于写case语句，如果当前Channel有缓冲可写，则执行。
func fibonacci(c, quit chan int) {
	println("B")
	x, y := 1, 2
	for {
		select {
		//写
		case c <- x:
			x, y = y, x+y
			//读
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
	//关闭 c管道
	close(c)
}

func main() {
	//非缓存需要及时取出来
	c := make(chan int)
	quit := make(chan int)

	go func() {
		fmt.Println("A")
		for i := 0; i < 5; i++ {
			//读
			fmt.Println(<-c)
		}
		//读完之后才来执行这步
		quit <- 0 //这个注释掉会造成死锁
	}()
	//写入数据   读写先后顺序？ ---产生了这个疑问
	//调度器？
	fibonacci(c, quit)
}

//为什么channel需要在goroutine执行？
//goroutine是Go语言的基本调度单位，而channels则是它们之间的通信机制。
// 操作符<-用来指定管道的方向，发送或接收。如果未指定方向，则为双向管道。
//另外一个问题？从channel读取的数据会减去吗 先进先出
//通过Goroutine能够让你的程序以异步的方式运行，而不需要担心一个函数导致程序中断，因此Go语言也非常地适合网络服务。
//一个程序如果以异步的形式运行 就不怕一个函数发生错误导致程序发生中断

//Goroutine中使用recover
//应用场景，如果某个goroutine panic了，而且这个goroutine里面没有捕获(recover)，那么整个进程就会挂掉。所以，好的习惯是每当go产生一个goroutine，就需要写下recover。

//Goroutine 栗子（等待所有任务退出主程序再退出）

//channel，管道、队列，先进先出，用来异步传递数据。
// channel加上goroutine，就形成了一种既简单又强大的请求处理模型，使高并发和线程同步之间代码的编写变得异常简单。

//线程安全，多个goroutine同时访问，不需要加锁。

//问题三 channel与异步？
//Channel主要用在go routine之间作为异步通信工具。
//简单说我们可以把Channel理解为一个队列，有人负责往里面写，有人负责从里面读，Channel会保证读和写操作的时序性和原子性，先写入的数据一定先读出来，而且一次读写都是一个完整的数据类型。

//首先一门语言在web领域都是需要异步的
//而go语言就是用goroutine解决异步，channel作为通信管道
//defer也是种异步啊
//goroutine为什么可以作为异步 按编译来说 我执行到这串代码是分在协程执行（它的执行方式是异步的就以上述的例子可能先打印出B)并且有多个协程（并行）

//判断是否关闭channel
//v, ok := <- ch
