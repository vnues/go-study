package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了确保同步交换数据的机制。（这是除了 atomic 和 mutex 之外的第三种处理竞态资源的方式
*/
// 正如官方所言，goroutine 是一个轻量级的执行单元，相比线程开销更小，完全由 Go 语言负责调度
/*
Channel 是 Go 中为 goroutine 提供的一种通信机制，借助于 channel 不同的 goroutine 之间可以相互通信。channel 是有类型的，而且有方向，可以把 channel 类比成 unix 中的 pipe。
Go 通过 <- 操作符来实现 channel 的写和读，send value <- 在 channel 右侧，receive value <- 在左侧，receive value 不赋值给任何变量是合法的。
*/
// 如果不指定容量，默认通道的容量是0，这种通道也成为非缓冲通道。
// 无论发送操作还是接受操作一开始就是阻塞的，只有配对的操作出现才会开始执行。
// 引起死锁的场景
/*
func main(){场景1：一个通道在一个go协程读写
func main() {
	c:=make(chan int)
	c<-666
	<-c
}
场景二：go程开启之前使用通道
func main() {
	c:=make(chan int)
	c<-666
	go func() {
		<-c
	}()
}
场景三：通道1中调用了通道2，通道2中调用通道1
func main() {
	c1,c2:=make(chan int),make(chan int)
	go func() {
		for  {
			select{
				case <-c1:
					c2<-10
			}
		}
	}()
	for  {
		select{
		case <-c2:
			c1<-10
		}
	}
}

*/

// WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
// Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。
// https://studygolang.com/articles/12972
var wg sync.WaitGroup

// Channel 完整的类型是 "chan variable"
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 接球
		// 1. 阻塞等待接球，如果通道关闭，ok返回false
		ball, ok := <-court
		if !ok {
			fmt.Printf("channel already closed! Player %s won\n", name)
			return
		}

		random := rand.Intn(100)
		if random%13 == 0 {
			fmt.Printf("Player %s Lose\n", name)
			// 关闭通道
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 2. 发球，阻塞等待对方接球
		court <- ball
	}
}

func main() {
	// 这样写造成 deadlock
	// ch := make(chan int)
	// ch <- 1
	// x := <- ch
	// ch <- 2
	// fmt.Println(ch)
	// fmt.Println(x)
	// 换句话话说，这里的 time.sleep 提供的是一种调度机制，这也是 Go 中 channel 存在的目的：负责消息传递和调度。
	// channel 是有类型的，而且有方向
	// Channel 最重要的作用就是传递消息。
	// c := make(chan int)
	// channel通信需要在goroutine进行
	// 创建一个goroutine
	// 执行
	/*go func(){
		fmt.Println("goroutine message")
		c <- 1
	}()*/
	// 只能执行一次读 因为写入就需要读 这样才不会造成阻塞 第二次读读时候你得写入
	// <- c
	// 读
	// fmt.Println(<- c)
	// 写入需要在goroutine执行
	// go func(){
	//	fmt.Println("goroutine message")
	//	c <- 2
	// }()
	// fmt.Println(<- c)
	// 不能这样写
	// x:=<- c
	// fmt.Println(c) //引用类型
	// fmt.Println("main function message")

	/*接收操作先发生*/

	// c :=make(chan int)
	//	//fmt.Println(<- c)
	//	//go func(){
	//	//	c <- 111
	//	//}()
	//	//fmt.Println(<- c)
	//	//time.Sleep(1)

	// 两个 player 打网球，即生产者和消费者模式（互为生产者和消费者）

	// wg.Add(2)
	//
	// // 1. 创建一个无缓冲的通道
	// // Channel 完整的类型是 "chan variable"
	// court := make(chan int)
	//
	// // 2. 创建两个 goroutine
	// go player("zhangsan", court)
	// go player("lisi", court)
	//
	// // 3. 发球：向通道发送数据，阻塞等待通道对端接收
	// court <- 1
	//
	// // 4. 等待输家出现
	// 相当于time.Sleep
	// wg.Wait()

	c := make(chan bool, 100)
	// 可以先塞进去100个因为它是可缓存的 如果是非缓存的那么它需要同步的读写
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
		// fmt.Printf("i is %d"+"c is %v ",i,<-c)
	}

}

// 高并发秒杀 谁先进来谁就赢 channel就是这个管道排队  --这个服务器能不能承受--并行？
// channel有阻塞作用
// 也就是channel有等待任务结束的功能

/*
chan T          // 可以接收和发送类型为 T 的数据

chan <- float64  // 发送：只可以用来发送 float64 类型的数据

<-chan int      // 接收：只可以用来接收 int 类型的数据
*/

// channel依居于线程下  -- 不然怎么运行
