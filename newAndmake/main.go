package main

import "fmt"

// 变量的声明我们可以通过var关键字，然后就可以在程序中使用。当我们不指定变量的默认值时，这些变量的默认值是他们的零值，
// 比如int类型的零值是0,string类型的零值是""，引用类型的零值是nil。
func main() {
	// 声明变量
	var s string        // 没有赋值
	fmt.Printf("%s", s) // 打印出来是空
	// 引用声明

	// var i *int //指针是int类型的？还是值是int类型
	// 值是int类型 而且这个类型是int引用类型
	// *i=88
	// fmt.Printf("%d",i)//报错

	// 用new分配内存
	// new是引用声明 new(int)==>*int
	// 它只接受一个参数，这个参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针。同时请注意它同时把分配的内存置为零，也就是类型的零值
	// 这就是new，它返回的永远是类型的指针，指向分配类型的内存地址
	var n *int
	n = new(int)
	*n = 88
	println(*n)

	i := new(int)
	*i = 881
	println(*i)
	// make声明
	// make也是用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，
	// 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	// 注意，因为这三种类型是引用类型，所以必须得初始化，但是不是置为零值，这个和new是不一样的。
	/*
	   从函数声明中可以看到，返回的还是该类型。
	   func make(t Type, size ...IntegerType) Type
	*/
	var m = make(map[string]int)
	m["alice"] = 31
	m["charlie"] = 34
	// 返回的是该类型
	// slice map channel这三个是引用类型
	println(m) // 0xc00007ae88
	println(m["alice"])

	var r = new(map[string]int)
	println(r) // 0xc000072e70

}
