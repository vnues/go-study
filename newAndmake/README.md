## new和make
>Go语言中new和make是内建的两个函数，主要用来创建分配类型内存。在我们定义生成变量的时候，可能会觉得有点迷惑，其实他们的规则很简单，下面我们就通过一些示例说明他们的区别和使用。

#### 变量的声明
```go
var i int
var s string
```

>变量的声明我们可以通过var关键字，然后就可以在程序中使用。当我们不指定变量的默认值时，这些变量的默认值是他们的零值，比如int类型的零值是0,string类型的零值是""，引用类型的零值是nil。对于例子中的两种类型的声明，我们可以直接使用，对其进行赋值输出。但是如果我们换成引用类型呢？

```go
package main

import (
	"fmt"
)

func main() {
	var i *int
	*i=10
	fmt.Println(*i)

}
//这个例子会打印出什么？0还是10?。以上全错，运行的时候会painc，原因如下
//anic: runtime error: invalid memory address or nil pointer dereference
```

从这个提示中可以看出，对于引用类型的变量，我们不光要声明它，还要为它分配内容空间，否则我们的值放在哪里去呢？这就是上面错误提示的原因。

对于值类型的声明不需要，是因为已经默认帮我们分配好了。

要分配内存，就引出来今天的**new**和**make**。

## new(new(T) 返回的是 T 的指针)

```go
//new 声明
 i :=new(int)
 r := new(map[string]int)
 println(r)//0xc000072e70

 /*
 现在再运行程序，完美PASS，打印10。现在让我们看下new这个内置的函数。
 func new(Type) *Type
 */
 //它只接受一个参数，这个参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针。
 //同时请注意它同时把分配的内存置为零，也就是类型的零值。
```

## make(make 只能用于 slice,map,channel)

>make也是用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。注意，因为这三种类型是引用类型，所以必须得初始化，但是,不是置为零值，这个和new是不一样的。

- 注意slice map channel这三个是引用类型
- make(T, args) 返回的值通过函数传递参数之后可以直接修改，即 map，slice，channel 通过函数穿参之后在函数内部修改将影响函数外部的值。
```go
/*
从函数声明中可以看到，返回的还是该类型。
func make(t Type, size ...IntegerType) Type
*/
ages := make(map[string]int)
println(ages)//0xc00007ae88 //引用类型
//map字面值的语法创建map，同时还可以指定一些最初的key/value：
ages := map[string]int{
    "alice":   31,yy
    "charlie": 34,
}

```
 
## 其实new不常用

所以有new这个内置函数，可以给我们分配一块内存让我们使用，但是现实的编码中，它是不常用的。我们通常都是采用短语句声明以及结构体的字面量达到我们的目的，比如：

```go
i:=0
```
**make函数是无可替代的，我们在使用slice、map以及channel的时候，还是要使用make进行初始化，然后才才可以对他们进行操作。**

