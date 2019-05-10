package main

import "fmt"

func doWork(id int,c chan int,d chan bool) {

	for  n := range c{
		//接收数据
		fmt.Printf("worker %d received %c\n", id, n)
		go func(){
			d<-true
		}()
		//这样写是并发？应该是并行执行 就是可能 c已经接收再写入再执行这步它还没来接收又写入数据就造成死锁
		//所以解决方法就是仔开一个routine你有数据我再接收 这样不是并行的就不用等对方什么的
		//d<-true
	}

}

type worker struct {
	in chan int
	done chan bool
}


func  createWorker(id int) worker {
	//非缓存通道
	w :=worker{
		in:make(chan int),
		done:make(chan bool),
	}
	go doWork(id,w.in,w.done)
	return w
}

func chanDemo(){
	var workers [10] worker
	//创建10个goroutine
	for i:=0;i<10;i++{
		//创建10个worker结构体 10个in管道
		workers[i]=createWorker(i)
	}
	for i:=0;i<10;i++{
		//写入数据
		//1-》10个管道进行通信
		workers[i].in <- 'a'+i
		//通过第二个channel阻塞来控制是否完成
		//但是这种是顺序的
		//<-workers[i].done
	}
	for i:=0;i<10;i++{
		workers[i].in <- 'A'+i
		//<-workers[i].done
	}
	//有数据就拿--显然这个通信快于这步的
	//等他们全部写入完再读取
	//循环
	for _,worker := range workers{
		<-worker.done

		<-worker.done
	}
}

func main(){
	//问题一他们为什么会乱顺序
	//有10个管道 开了10个goroutine相当于main（goroutine）在跟这10个goroutine通信
	//而且并行的
  chanDemo()
  //time.Sleep(1)
}
