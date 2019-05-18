## 浅谈slice、byte、string、rune


### slice(引用类型)
> 在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice

`slice`并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层`array`，只是不需要长度。slice的声明也可以像array一样

声明：
```go
//声明一个保存 int 的 slice
var iSlice []int

//声明一个长度为 10 的 int 数组
var iArray [10]int

//还有一种声明的方法是使用 make() 函数，如下
slice1 := make([]int, 5, 10)
//用 make() 函数创建的时候有三个参数，make(type, len[, cap]) ，依次是类型、长度、容量。
```
 slice底层代码
```go
type IntSlice struct {
    ptr      *int
    len, cap int

```
**注意`slice`和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用`...`自动计算长度，而声明`slice`时，方括号内没有任何字符。**

slice有一些简便的操作

 - `slice`的默认开始位置是0，`ar[:n]`等价于`ar[0:n]`
 - `slice`的第二个序列默认是数组的长度，`ar[n:]`等价于`ar[n:len(ar)]`
 - 如果从一个数组里面直接获取`slice`，可以这样`ar[:]`，因为默认第一个序列是0，第二个是数组的长度，即等价于`ar[0:len(ar)]`
 
 
 ### string）
 >，Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string
 
 **注意go语言的string类型与其他语言的string类型可能不一样，比如看以下代码:**
 
 ```go
     s3 :="a我爱中国"
	for _,ss:=range s3{
		/*
		 97
		25105
		29233
		20013
		22269
		*/
		fmt.Println(ss)
	}
```
**很神奇的一件事就是string能被像数组一样循环，我们来看看string的底层源码**
```go
	 type stringStruct struct {
	    str unsafe.Pointer
	    len int
	}
//实际上string是个结构体指针指向[]byte，也就是说string底层是byte数组，所以可以被循环，也就是可以通过len方法知道具体长度（中英文长度可能不一样，稍后再说）

```
-在Go中字符串是不可变的,例如下面的代码编译时会报错：cannot assign to s[0]
```go
var s string = "hello"
s[0] = 'c
```
但如果真的想要修改怎么办呢？下面的代码可以实现：
```go
//其实这样的修改就是拷贝
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
//"c"会报错 因为这是字符串而不是字符字符用单引号表示''
c[0] = 'c' //这样的复制很奇怪 因为byte是uint8也就是数值 隐式转化？我的猜测
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```
- string的结构体指针指向的是byte数组
```go
//utf-8中英文占一个字节，中文是三个，而字符串底层是byte数组,s[0]读出来自然是乱码
//还有我感觉string就是复合类型 结构体对象并且指针指向byte数组
func testString() {
    s := "我是中国人"
    // 错误用法：读取“我”
    fmt.Println("%c",s[0]) // 输出：æ
}
```

### byte与rune
来看个有趣的现象：
```go
 first :="first"
  fmt.Println([]rune(first))  //[102 105 114 115 116]
  fmt.Println([]byte(first))  //[102 105 114 115 116]
 china :="你好，中国"
   fmt.Println([]rune(china)) //[20320 22909 65292 20013 22269]
   fmt.Println([]byte(china)) //[228 189 160 229 165 189 239 188 140 228 184 173 229 155 189]
```
>byte类型对应着uint8，说到byte就不得不说中文情况下，字符是占多少个字节，因为前面说了字符串底层是指向byte数组的，首先go的编码形式是utf-8编码
>而rune是int32的别名，代表字符的Unicode编码，采用4个字节存储，将string转成rune就意味着任何一个字符都用4个字节来存储其unicode值，这样每次遍历的时候返回的就是unicode值，而不再是字节了，这样就可以解决乱码问题了
- UTF-8编码：一个英文字符等于**一个字节**，一个中文（含繁体）等于**三个字节**。中文标点占三个字节，英文标点占一个字节
- Unicode：计算机是二进制的，字符最终也是转换成二进制保存起来的。字符集就是定义字符对应的数值。 Unicode是一个字符集，为每个字符规定一个用来表示该字符的数字，但是并没有规定该数字的二进制保存方式，utf8规定了对于unicode值的二进制保存方式
 **大白话就是一组数字对应一个字符**
- go用rune来表示unicode码点，也就是rune类型字符都是占**四个字节**不管中英文


### 为什么slice打印出来的是字符而不是byte数组

```go
  s3 :="a我爱中国"
	fmt.Println(s3[:]) //a我爱中国
	fmt.Println(s3[:2])//a�
	fmt.Println(s3[:4])//a我 0123   以字节来算的
```
//go内部这样转化为string([]byte(s))？