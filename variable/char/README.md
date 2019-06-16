# 字符类型(char)-区别好字符串类型string

## 基本介绍
> Golang 中没有专门的字符类型，如果要存储单个字符(字母)，一般使用 byte 来保存。字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。也 就是说对于传统的字符串是由字符组成的，而 Go 的字符串不同，它是由字节组成的。

```go
import (
	"fmt"
)

//演示golang中字符类型使用
func main() {
	
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//当我们直接输出byte值，就是输出了的对应的字符的码值
	// 'a' ==> 
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '北' //overflpackage main
                                
	var c3 int = '北' //overflow溢出
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3)
	}
```
对上面代码说明
- 如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]直接可以保存到 byte
- 如果我们保存的字符对应码值大于 255,这时我们可以考虑使用 int 类型保存
```go
var c3 byte = '北' //overflow溢出 所以说中文用rune类型存中文占utf-8占三个字节 go用int32 rune四个字节存
```
> Golang中是以utf-8形式编码的,英文占一个，中文占三个,而一个带有中英文字符串是以byte数组形式存的，英文占一个数组元素，中文三个(如果是用byte存就会溢出也就是三个数组元素才能形成中文)，所以你想打印这个字符串的其中一个中文需要转成rune类型才行
```go
package main


import "fmt"

func main(){
	var c = "I love 中国"
	fmt.Println(c[2]) //108
	fmt.Println(string(c[2])) // l
	fmt.Println(string(c[7])) // 乱码 三个字节才能形成中文
	fmt.Println([]byte(c)) // [73 32 108 111 118 101 32 228 184 173 229 155 189]
	// 转化成rune类型数组
	c2:=[]rune(c)
	fmt.Println(c2) // [73 32 108 111 118 101 32 20013 22269]
	fmt.Println(string(c2[8])) // 中
	fmt.Println(string(c2[8])) // 国
}

```

- 如果我们需要安装字符的方式输出，这时我们需要格式化输出，即 fmt.Printf(“%c”, c1)..
- 整型类型包括int unit rune byte(unit8等价)，而string类型实际是用byte类型保存的,所以说go语言string类型特殊
```go
// 官方源代码这样定义byte
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8
```
对上面的代码说明：
- 开题抛出byte类型存字符串 但是只是针对于英文情况下一个英文对应一个byte但是中文不一定，go以utf-8形式编码，而中文在utf-8是占三个字节，如果以byte类型存中文显然会溢出，所以为了解决上述情况,go用rune类型存
- 字符串底层就是Byte数组
- []byte方法和[]rune方法是将字符串转化成对应类型的slice(slice是引用类型，引用数组)，string方法是转化为字符串


## 字符类型使用细节
- 字符常量是用单引号('')括起来的单个字符。例如:var c1 byte = 'a' var c2 int = '中' var c3 byte = '9'

- Go 中允许使用转义字符 '\’来将其后的字符转变为特殊字符型常量。例如:var c3 char = ‘\n’ // '\n'表示换行符

- 在 Go 中，字符(注意是字符不是字符串 字符串直接打印还是显示字符串，但是打印单个字符就是整型了)的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值

- 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的 unicode 字符

```go
	//可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
	var c4 int = 22269 // 22269 -> '国' 120->'x'
	fmt.Printf("c4=%c\n", c4)
```

- 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码.
```go
//字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
	var n1 = 10 + 'a' //  10 + 97 = 107
	fmt.Println("n1=", n1)
```

## 字符类型本质探讨

- 字符型 存储到 计算机中，需要将字符对应的码值(整数)找出来 存储:字符--->对应码值---->二进制-->存储
  读取:二进制----> 码值 ----> 字符 --> 读取
  
- 字符和码值的对应关系是通过字符编码表决定的(是规定好)

- Go语言的编码都统一成了utf-8。非常的方便，很统一，再也没有编码乱码的困扰了