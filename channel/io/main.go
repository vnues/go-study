package main

import (
	"fmt"
	"time"
)

func main(){
	//io 操作就是输入和输出
	//golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
	//注意到 select 的代码形式和 switch 非常相似， 不过 select 的 case 里的操作语句只能是【IO 操作】 。
	//除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
	//
	//除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
	//
	//如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
	//
	//如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞

	//为请求设置超时时间
	//c :=make(chan int)
	//c<-1
	timeout := time.After(5 * time.Second)
	/*
	func After(d Duration) <-chan Time {
		return NewTimer(d).C
	}
	*/
	fmt.Println(timeout)
	for {
		select {
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}

	//goroutine怎么保证退出
}