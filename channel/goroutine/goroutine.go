package main

import (
	"fmt"
	"time"
)
//不能在函数再定义其他非匿名函数 --所以函数传参的必要
func worker(id int,c chan int) {
	for {
		n, ok := <-c
		if !ok {
			fmt.Println("数据接收终止")
			break
		}
		fmt.Printf("worker %d received %d\n", id, n)
	}
	//效果类似的 意思就是等到d发完就跳出来
	//for  n := range c{
	//	fmt.Printf("worker %d received %d\n", id, n)
	//}

}
func main(){
	  //注意channel是goroutine之间的通道，也就是起码至少两个goroutine
	 // A-channel-B
	 //这就是为什么需要在一个goroutine去读或者写
	 //以下这种写法会造成死锁，因为只有一个goroutine对象没有另外一个怎么建立通道
	 //c := make(chan int)
	 //fmt.Println( <- c)
	 //有两个goroutine进行通信 --但是我变化下我先把读的函数写在前面，比如：
	//c := make(chan int)
	//A
	// go func(){
	//	 fmt.Println("读取")
	// 	 fmt.Println(<-c)
	// }()
	 //fmt.Println("写入")
	//B
	// c<-11
	//C
	//go func(){
	//	fmt.Println("读取")
	//	fmt.Println(<-c)
	//}()
	//fmt.Println("写入")
	// c<-22
	 /*
	  读取
	  写入
	  11
	 也可能
	 写入
	 读取
	 11

	 结果是以上的 我们不能按照代码同步的逻辑去分析，而是按照通信的，而且goroutine是异步的
	  fmt.Println(<-c)是读取，发现B还没有发信息，我阻塞了--等它发信息过来我再执行
	 */
	 //升级 我们写个死循环不断接收数据

    // func we(){
    // 	println("11")
	//}

	 c := make(chan int)
	 go worker(0,c)

     c <- 11
     c <- 22
    // c <- 'a' //a--ascii 97?
     //c <- 'a的' 会报错
     //增加close
     close(c)
     time.Sleep(20)

}

//调度就是控制这个队列进出（只有我1出去2才能出去 按队列顺序）

//建立通信只要使用c就是这两个goroutine就建立通信