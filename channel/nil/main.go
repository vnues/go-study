package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			//写
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	//一直在读 如果没有close掉channel
	for n := range c {
		//read
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

//channel是引用传值
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	//怎么会一直执行下去？
	//因为有写有读 一直在通信下去
	//为什么会打印
	//fmt.Println(nil <- 1)
	for {

		select {
		//读
		//w<n与fmt.Printf("Worker %d received %d\n",id, n)建立了通信  这样理解 那里会有阻塞？
		//只是进去case以后还有一个通道通信 在线程中相当于并发了 我需要你这个走完才可以走
		//并发是这样理解读吗
		case n := <-c1:
			w <- n ////这样不好，收一个数之后，后面的操作又会阻塞,最后的出的结果是由顺序的，一个任务结束之后才可以进行下一个
			//w<-1
		case n := <-c2:
			w <- n

		}
	}
}

//定义只读的channel
//read_only := make (<-chan int)
//定义只写的channel

//rite_only := make (chan<- int)

//可同时读写

//read_write := make (chan int)

//当activeWork是nil时，这里一直是block(阻塞)
//nil channel 在case的应用是nil接收变量 会造成阻塞而不是panic

/*
当case上读一个通道时，如果这个通道是nil，则该case永远阻塞。这个功能有1个妙用，select通常处理的是多个通道，当某个读通道关闭了，但不想select再继续关注此case，继续处理其他case，把该通道设置为nil即可。
下面是一个合并程序等待两个输入通道都关闭后才退出的例子，就使用了这个特性。
*/

/*
1、select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

2、select中的case语句必须是一个channel操作，select中的default子句总是可运行的。

3、如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。

4、如果没有可运行的case语句，且有default语句，那么就会执行default的动作。

5、如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

*/
