package main

import "fmt"

// 学习一门语言要知道他的编码形式以哪种形式很重要
// Go 语言采用 UTF-8 编码，这种编码方式与 ASCII 编码兼容，只不过 ASCII 编码只需 1 个字节，而 UTF-8 需要 1-4 个字节表示一个符号。
// 字符串类型声明 var s string
// 整数类型 int8 int32
// 字符串没有 string8...什么的
// 但是字符串类型比较和运算还是得转化成编码形式的 也就是归根结底的int32
// []...以数组形式的方法
// rune理解为 一个 可以表示unicode 编码的值int 的值
// go本身是utf-8编码的，char不能满足要求，所有go就有rune这个类型。
// rune 是 Go 的内置数据类型，是 int32 的别名，表示 Go 中的 Unicode 代码点。用 rune variable，开发人员就不必关心代码点占用几个字节了。
// 所以字符串类型就是用rune表示

// go语言循环字符串输出的结果是rune类型
// https://juejin.im/post/5c1a2db5f265da61682b52f5
// rune解决乱码问题
// 就是将原来分散的三个字节（中文）聚集存起来
// rune是int32的别名，代表字符的Unicode编码，采用4个字节存储，将string转成rune就意味着任何一个字符都用4个字节来存储其unicode值，这样每次遍历的时候返回的就是unicode值，而不再是字节了，这样就可以解决乱码问题了
// 计算机是二进制的，字符最终也是转换成二进制保存起来的。字符集就是定义字符对应的数值。
// Unicode是一个字符集，为每个字符规定一个用来表示该字符的数字，但是并没有规定该数字的二进制保存方式，utf8规定了对于unicode值的二进制保存方式。
func main() {

	// s := "hello 世界"
	// s = "hello 世界qqq"
	// fmt.Println(s)
	// c := "界"
	// fmt.Printf("%T ", c) //string类型
	// //unicode形式
	// //这样走的
	// //字符串的前提下 utf-8->unicode->rune(四个字节)
	// //因为是1236...这样开始是按utf-8编码再转化成unicode
	// for i, ch := range s {
	//	//ch是rune类型 但是s[i]是uint8类型 这个要区别好
	//	//ch是int32类型的 相当于rune类型
	//	//字符串转化成编码形式的话就是int32位类型
	//	fmt.Println(i, ch)
	//	fmt.Printf("%T ", ch)
	// }
	// for i, ch := range []rune(s) {
	//	//ch是int32类型的 相当于rune类型
	//	//字符串转化成编码形式的话就是int32位类型
	//	fmt.Println(i, ch)
	//	fmt.Printf("%T ", ch)
	// }
	// //utf-8形式
	// for i, ch := range []byte(s) {
	//	fmt.Println(i, ch)
	// }
	// fmt.Println("------------------------------------")
	// s1 := "hello 世界"
	// for i, ch := range s1 {
	//	fmt.Println(i, ch)
	// }
	// fmt.Println("------------------------------------")
	// for i, ch := range []byte(s1) {
	//	fmt.Println(i, ch)
	// }
	// fmt.Println("------------------------------------")
	// m := make(map[byte]int)
	// fmt.Println(m)
	// fmt.Println("------------------------------------")
	// for i, ch := range []byte(s) {
	//	m[ch] = i
	//	fmt.Println(m[ch])
	// }
	// fmt.Println(m)
	// fmt.Println("------------------------------------")
	// lan := "我爱中国"
	// for _, ch := range []rune(lan) {
	//	//fmt.Println(i, ch)
	//	//rune是unicode编码的表现
	//	println(ch)
	//	fmt.Printf("%c ", ch)
	// }
	// slice 必须指定长度需要的话在用append追加 map可用
	// var monsters []map[string]string
	// monsters= make([]map[string]string,2)
	// if(monsters[0]==nil){
	//
	//	monsters[0]=make(map[string]string)
	//	monsters[0]["name"]="林晓珊"
	//	monsters[0]["age"]="38"
	//	monsters[0]["height"]="118"
	// }
	// fmt.Println(monsters)
	// var m map[string]string
	// m =make(map[string]string,2)
	// m["no1"]="宋江"
	// m["no2"]="吴用"
	// // map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动 态的增长 键值对(key-value)
	// // m["no3"]="武松" // map[no1:宋江 no2:吴用 no3:武松]
	// fmt.Println(m) // map[no1:宋江 no2:吴用]

	// A slice B数组

	var A [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(A)
	B := A[0:3]
	fmt.Println(B)
	fmt.Println(&A[0])                // 0xc0000a0000
	fmt.Println(&B[0])                // 0xc0000a0000 肯定一样啊 因为B[0]这个操作就是他里面strcut指向的
	fmt.Println(&B)                   // &[1 2 3]
	fmt.Printf("A指针是 %p\n", &A)       // 0xc00009c000
	fmt.Printf("B[0]指针是 %p\n", &B[0]) //  0xc00009c000
	fmt.Printf("B指针是 %p", &B)         // 0xc0000a6000
	// 指向这个slice的内存指针跟A确实不一样 那肯定啊 内存得有帮slice开辟内存的吧
}

// go参数的传递就只有一种 值传递 可通过指针拷贝（实际就是值传递）到达引用传递的效果

// 如果重复调用某函数 而且传参数有一个是固定的 那么就用闭bao会优雅 f2('jpg',22),f2('jpg',11)

// go数组是直接指向内存空间 而不是指向地址所以数组是值类型
// slice是引用类型 竟然是引用那么就得有对应得引用对象 数组
// slice定义完后还不能使用需要让其引用到一个数组或者make一个空间
// 指针指向内存

// 指针放在堆里面可以共享
// 什么是引用类型 -- 可以理解成他们执行的是变量内存 比如说slice指向的是数组内存
// 在golang中只有三种引用类型它们分别是切片slice、字典map、管道channel。其它的全部是值类型，引用类型可以简单的理解为指针类型，它们都是通过make完成初始化
// map 是 key-value 数据结构，又称为字段或者关联数组。类似其它编程语言的集合，
// 在编程中是经常使用到
// map跟我们javascript语言的json对象很相似 但是map存放的数据类型不能不同一定要一样 所以通常用map+struct->js对象

// 引用类型跟指针做好区别 虽然实际是一样的 都是指向那块变量内存
// 但是语言区分类型是因为你赋值或者参数传递是以拷贝还是指向内存空间的 注意 指针也是指向内存空间的
// 注意在golang中只有三种引用类型它们分别是切片slice、字典map、管道channel。其它的全部是值类型，引用类型可以简单的理解为指针类型，它们都是通过make完成初始化
// 值类型说明传递或者赋值就是拷贝
// 你可能会很奇怪为什么我打印的map类型不是一个指针（地址），因为引用类型又不是指针，map里面有做了对应的底层处理让他可以指向变量内存
// 注意引用类型指向内存变量 值类型也是指向内存变量 （不指向他怎么拿到）？对
// 不同的是引用类型指向公共的 值都是自己开辟自己的 -- 好像就是放在栈还是放在堆

// --------------------------------------------- 看这里
/*
再总结一遍 golang的参数传递只有值传递也就是拷贝 即使你传个指针 也是拷贝了指针的

值类型：int、float、bool和string这些类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
(对啊这个变量指向内存 不然怎么拿到值，而我们所说的值类型只是说在一系列赋值或者参数传递是拷贝啊还是指向同一个内存这才是他们的区别 变量永远都是指向内存的 这对于你画内存析图很重要)
(很重要的一点 以后记得写在笔记里)
值类型的变量的值存储在栈中。当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 &i 获取变量 i 的内存地址。  值拷贝

引用类型：特指slice、map、channel这三种预定义类型。引用类型拥有更复杂的存储结构:(1)分配内存
(2)初始化一系列属性等一个引用类型的变量r1存储的是r1的值所在的内存地址（数字）
或内存地址中第一个字所在的位置，这个内存地址被称之为指针，这个指针实际上也被存在另外的某一个字中。

两者的主要区别：拷贝操作和函数传参。

来看下slice的数据结构
slice 从底层来说，其实就是一个数据结构(struct 结构体) type slice struct {
ptr *[2]int
len int cap int

}

slice底层确实存了一个指针所以指向内存 但是他本身自己也是个内存也就是这样的工作形式的 A slice 这个变量你得有内存去存它吧 而它这个变量里面又存了个指针 指向B数组
我们要操作的就是这个B数组 其实本身A内存我们不会过多的关注

*/
