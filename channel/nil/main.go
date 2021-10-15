package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1、select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

2、select中的case语句必须是一个channel操作，select中的default子句总是可运行的。

3、如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。

4、如果没有可运行的case语句，且有default语句，那么就会执行default的动作。

5、如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
*/

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 相当于延时
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			// 写
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	// 一直在读 如果没有close掉channel
	for n := range c {
		// read
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

// channel是引用传值
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	for {
		select {
		case n := <-c1:
			w <- n
		case n := <-c2:
			w <- n
		}
	}
}
