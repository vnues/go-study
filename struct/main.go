package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode //treeNode类型
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

//go指针也很方便直接xx.xx的形式也可以访问属性
//引用者
func (node *treeNode) setValue(value int) {
	//value int传递过来的参数
	node.value = value
}

//go的工厂函数
//工厂函数一般返回一个地址
//https://juejin.im/post/5bdbcc08f265da61561eb493
//理解有点困难 工厂模式
//结论就是一个函数内局部变量，不管是不是动态new出来的，它会被分配在堆还是栈，是由编译器做逃逸分析之后做出的决定。
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 2}
	root.left = &treeNode{value: 3}
	//数组形式
	nodes := []treeNode{
		{value: 3},
		{},
	}
	fmt.Println(root)
	fmt.Println(nodes)
	//方法接收者有这样一个处理就是不管接收者是值接收者或者是址接收者都会做转化 如果接收者是地址 但是方法需要的是值 地址就会取出值给他 值接收者反过来也是这样的
	//值引用
	root.print()
	//址传递
	(&root).setValue(20)
	root.print()
	root.left.setValue(4)
	root.left.print()
}

// 一、C语言中返回函数中局部变量值和指针
// (1) 在C语言中，一个函数可以直接返回函数中定义的局部变量，其实在函数返回后，局部变量是被系统自动回收的，因为局部变量是分配在栈空间，那为什么还可以返回局部变量，其实这里返回的是局部变量的副本（拷贝）。

// (2) 函数返回局部变量地址：局部变量内存分配在栈空间，因为函数返回后，系统自动回收了函数里定义的局部变量，所以运行时去访问一个被系统回收后的地址空间，一定就会发生段错误，这是C/C++语言的特点。内存空间分配在堆中即可。
// 二、GO函数中返回变量，指针
// 示例代码：

// package main

// import "fmt"

// func fun() *int {    //int类型指针函数
//     var tmp := 1
//     return &tmp      //返回局部变量tmp的地址
// }

// func main() {
//     var p *int
//     p = fun()
//     fmt.Printf("%d\n", *p) //这里不会像C，报错段错误提示，而是成功返回变量V的值1
// }
// 参考go FAQ里面的一段话：

// How do I know whether a variable is allocated on the heap or the stack?

// From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.

// The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

// In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.
// 意思是说go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。所以不用担心会不会导致memory leak，因为GO语言有强大的垃圾回收机制。go语言声称这样可以释放程序员关于内存的使用限制，更多的让程序员关注于程序功能逻辑本身。
// 对于动态new出来的局部变量，go语言编译器也会根据是否有逃逸行为来决定是分配在堆还是栈，而不是直接分配在堆中。

//结构体就是其他语言的面向对象
//接口？

//方法和函数的区别
//对象行为 -->方法 但是go的对象方法定义是这样的
//根据参数来确定这个方法接受者是谁 也就是简单讲就是这个结构体的方法
// func (p *Person) GetName() string {
//     return p.Name
// }
// https://juejin.im/post/5be28fbfe51d4517ad11147e
