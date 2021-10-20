package main

import "fmt"

func main() {
	// var c2 byte ='北' // constant 21271 overflows byte
	var c3 int = '北'
	//fmt.Println(c2)
	fmt.Printf("c3的类型是 %T\n", c3)
	var c = "I love 中国"
	fmt.Println(c[2])         //108
	fmt.Println(string(c[2])) // l
	fmt.Println(string(c[7])) // 乱码 三个字节才能形成中文
	fmt.Println([]byte(c))    // [73 32 108 111 118 101 32 228 184 173 229 155 189]
	// 转化成rune类型数组
	c2 := []rune(c)
	fmt.Println(c2)            // [73 32 108 111 118 101 32 20013 22269]
	fmt.Println(string(c2[8])) // 中
	fmt.Println(string(c2[8])) // 国
}
