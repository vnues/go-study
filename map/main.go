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
