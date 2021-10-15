package main

// interface <- 实例化结构体 <- 结构体

//所有的类型都可以进行类型断言不然怎么叫类型断言
import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

// 打印
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

/*
Comma-ok 断言

Go 语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里 value 就是变量的值，ok 是一个 bool 类型，element 是 interface variable，T 是断言的类型。

如果 element 里面确实存储了 T 类型的数值，那么 ok 返回 true，否则返回 false。

package main

import (
    "fmt"
    "strconv"
)

type Element interface{}
type List [] Element

type Person struct {
    name string
    age int
}

// 定义了 String 方法，实现了 fmt.Stringer
func (p Person) String() string {
    return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
    list := make(List, 3)
    list[0] = 1 // an int
    list[1] = "Hello" // a string
    list[2] = Person{"Dennis", 70}

    for index, element := range list {
        if value, ok := element.(int); ok {
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        } else if value, ok := element.(string); ok {
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.(Person); ok {
            fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
        } else {
            fmt.Printf("list[%d] is of a different type\n", index)
        }
    }
}


*/
// switch类型断言
// 这里有一点需要强调的是：element.(type) 语法不能在 switch 外的任何逻辑里面使用，
// 如果你要在 switch 外面判断一个类型就使用 comma-ok。
// switch type： 已知或者未知的对象数据类型均可，b1.(type)必须配合switch来使用，不能单独执行此语句。
// 执行顺序是这样的 range循环 首先是list[0]->switch type->依次跟case比较 不过这个 switch value := element.(type)看起来怪怪的
// 语法有点特殊 v表示拿到这个循环item的值 像js可能我们拿到的v是用来做比较的 但是go中switch断言会自动帮我们比较的 switch value := element.(type)这步的是
// 判断item的类型并拿到他的值 我把value赋值去掉	switch  element.(type)这样类似我们js语法大概就这样 也不会报错
// 语法觉得怪异的拆出来看看就知道了
func main() {
	var a int = 2
	if value, ok := a.(int); ok {
		fmt.Println("进行类型断言value:%d\n", value)
	}
	// assertion type
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		// value :=1
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

// :=类型推断

// 小猴子类继承老猴子类的行为和方法 本身通过努力实现飞翔，那么我直接在小猴子类实现方法为什么通过实现接口尼因为猴子是不会飞的 你需要一个契机或者通道
// 这个契机或者通道就是interface 接口是对继承的扩展

// 空接口可以接收任何一个变量
