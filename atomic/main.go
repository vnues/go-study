package main



/*
有个问题就是并发是抢占资源的而 go的并发是非抢占式的？ 虽然是并发但是是非抢占式的
*/
/*
go的管理状态机制channel是主要的一种  还有（同步机制？）
并发是什么？
并发，指的是多个事情，在同一时间段内同时发生了。（段） 一个人干多件事（肯定需要时间长 所以是段）
并发的多个任务之间是互相抢占资源的。
并行的多个任务之间是不互相抢占资源的、
并行，指的是多个事情，在同一时间点上同时发生了（点） 多个人干多件事（不一定—））
*/

/*
但在Go 1.5以前调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
*/

/*
通信在go中应该属于异步的 各类编程语言中都有异步的概念，尤其是web领域的
Go里面的管理协程状态的主要机制就是通道通讯。这些我们上面的例子介绍过。这里还有一些管理状态的机制，下面我们看看多协程原子访问计数器的例子，这个功能是由sync/atomic包提供的函数来实现的
*/

/*
1.1、原子性
这一点，跟数据库事务的原子性概念差不多，即一个操作（有可能包含有多个子操作）要么全部执行（生效），要么全部都不执行（都不生效
*/


func main(){

}

//当你想利用多个线程并执行彼此独立的函数时会发生什么呢？这就是并发程序设计发挥作用的地方。
//go并发就是开启多个协程运行独立的函数（在这同个时间段）
//很明显我们的程序不可能老是等到a执行完再去执行B
//可以看出，两个 finder 是并发运行的。哪一个先找到矿石没有确定的顺序，当执行多次程序时，这个顺序并不总是相同的。
//就是多个任务在这时间段同时运行
//go的并发跟我们实际一些开发语言的并发有点差异
//假设多个任务在一个协程上运行 -- 》并发？ //并发是时间段概念上来讲
//假设多个任务在不同协程上（一个协程运行一个任务）并行吗？ //而并行是时间点
//但是go的协程可能会运行多个任务
//依靠多个协程（线程）来作为并发与并行的判断和区别吗？


//(高并发50亿条清楚 可能需要并行+并发  并行很少讲到 因为感觉就是默认)
//例子一 同一个时间段内假设这五分钟内 服务器收到50万条不同段请求 我们是单核cpu，这50万个请求在这时间段内完成



/*
go并发与并行的原理

并发(concurrency)：两个或两个以上的任务在一段时间内被执行。
我们不必care这些任务在某一个时间点是否是同时执行，
可能同时执行，也可能不是，我们只关心在一段时间内，
哪怕是很短的时间（一秒或者两秒）是否执行解决了两个或两个以上任务。



并行(parallellism)：两个或两个以上的任务在同一时刻被同时执行。

并发说的是逻辑上的概念，而并行，强调的是物理运行状态。并发“包含”并行
*/


//从这个例子我可以理解到并发的概念  并发（只要一个时间段内多个任务去处理 就是并发）
//但是并发归并发 还是离不开通信的

//高并发秒杀系统确实一个时间段内  比如说倒计时10分钟秒抢的


//那我之前做的那个竹笋模块属于并发吗 （并发在每个服务端语言都有这个概念吧 和实现的 不可能同步的 比如nodejs看似单线程--异步处理并发）
//而且我们本身写服务的时候不用考虑并发的 不可能同步的 （默认都是支持的 --标哥的意思）