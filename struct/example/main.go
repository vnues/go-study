package main

import (
	"fmt"
	"structExample/mock"
	"structExample/real"
	"time"
)

// Retriever TODO
// 定义啦这个接口就得实现这个接口
// 这个接口是抽象 它只告诉你它这个接口有哪些行为也就是方法
type Retriever interface {
	// 这个接口有get的方法
	Get(url string) string
	// Post(url string) string

}

// 声明这个接口解决实际问题的所需要的 把它归类成一个对象所有的行为？
// 然后我们再去调用它

func download(r Retriever) string {
	return r.Get("www.wgfvnues.mobi")
}

func main() {
	var r Retriever
	var test Retriever
	// 实例这个对象 后面肤质改变需要按照它这个接口格式要求
	// r=1
	// 接口变量r
	r = mock.Retriever{"this is mock"} // 实现这个接口
	test = mock.TestRetriever{"this is fake mock"}
	// 等价的  r := mock.Retriever{"this is mock"}
	fmt.Println(download(r))
	fmt.Println(download(test))

	// 类型断言判断
	// mockRetriever就是那个结构体对象 断言返回T
	// 类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)被称为断言类型，这里x表示一个接口的类型和T表示一个类型。一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。
	// 这里有两种可能。第一种，如果断言的类型T是一个具体类型，然后类型断言检查x的动态类型是否和T相同。如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。换句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，接下来这个操作会抛出panic。例如：
	mockRetriever, ok := r.(mock.Retriever)
	fmt.Println(mockRetriever)
	fmt.Println(ok)

	var myreal Retriever
	myreal = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Println(myreal.Get("http://www.hstc.edu.cn/"))

}

// 指针变量的类型
// 类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)被称为断言类型，这里x表示一个接口的类型和T表示一个类型。一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。
// 接口是放方法的吗？

// 接口的理解

// 就是我这个接口是使用者自己定义 随便定义
// 而我实现者不管你属于什么类型的接口我只要实现这个方法就行

// 那么接口的值是有什么尼？

// var w io.Writer // type-<nil>
// w = new(bytes.Buffer) // type-*bytes.Buffer
// w = nil // type-<nil>

// 接口是什么？一个接口包含两层意思：它是一个方法的集合，同样是一个类型。让我们首先关注接口作为方法的集合这一方面。

// 我们说这个 concrete type 实现了这个 interface 的所有方法。

// Type switch 就是利用 interface 存储不同 concrete type 的能力来实现的 assertion。

// r.(type)这样拿到类型

// 理解什么是接口变量

// var r Retriever //Retriever是定义的接口 r就是接口变量
// 接口变量有自己的指针
// 接口变量也是值传递
// 也就是传参数 接口变量是值传递？

// 接口变量自带指针？？？不明白

// https://segmentfault.com/a/1190000015591319

// 最重要的一点 就是接口就是一个struct吧

// https://blog.csdn.net/Ghost_leader/article/details/77126918

// 将对象复制给接口，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针

// 参数是一个局部变量
