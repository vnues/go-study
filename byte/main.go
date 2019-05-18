package main

import "fmt"

/*
字节（港澳台作位元组，英语：Byte），通常用作计算机信息计量单位，不分数据类型。[1][2] 一个字节代表八个比特（港澳台作位元，英语：Bit）。从历史的观点上，“字节”表示用于编码单个字符所需要的比特数量
。历史上字节长度曾基于硬件为1-48比特不等，最初通常使用6比特或9比特为一字节。今日事实标准以8比特作为一字节，因8为二进制整数。
*/

func main(){
  //在看到 go 字符串的时候， 偶然看到 []rune(s), 它可以将字符串转化成 unicode 码点。那么它和 []byte(s) 有什么区别呢？
  //[]rune是rune是数组 []rune(参数) 是rune数组方法
  first :="first"
  fmt.Println([]rune(first))  //[102 105 114 115 116]
  fmt.Println([]byte(first))  //[102 105 114 115 116]

  /*
  源码
  type byte = uint8
  // rune is an alias for int32 and is equivalent to int32 in all ways. It is
  // used, by convention, to distinguish character values from integer values.
  type rune = int32
  // iota is a predeclared identifier representing the untyped integer ordinal
  */

  //go的编码形式是utf-8编码
  //UTF-8编码：一个英文字符等于一个字节，一个中文（含繁体）等于三个字节。中文标点占三个字节，英文标点占一个字节

  china :="你好，中国"
  fmt.Println([]rune(china)) //[20320 22909 65292 20013 22269]
  fmt.Println([]byte(china)) //[228 189 160 229 165 189 239 188 140 228 184 173 229 155 189]
  //这里也可以很清晰的看出这里的中文字符串每个占三个字节， 区别也就一目了然了。
  //说道这里正好可以提一下 Go 语言切割中文字符串，Go 的字符串截取和切片是一样的 s [n:m] 左闭右开的原则，看一下例子
  /*
   为什么中文占三个字节也就是三个uint8 而rune是int32就是4个uint8? 让我们来看看下面的例子：
   很明显
  */
  s :="111截取中文"
  fmt.Println([]byte(s))//[230 136 170 229 143 150 228 184 173 230 150 135]
  //为什么可以用切片打印出中文
  fmt.Println(s[:2]) //�
  //底层会将中文转化成 []byte， 而不是 []rune
  fmt.Println(s[:3]) //截
  fmt.Println(string(s[0:3]))//截
  fmt.Println(string([]byte(s)))//截取中文
	//切片-->中文
  //那么该如何截取呢？这里就需要将中文利用 [] rune 转换成 unicode 码点， 再利用 string 转化回去， 来看一下例子。



  //Unicode:计算机是二进制的，字符最终也是转换成二进制保存起来的。字符集就是定义字符对应的数值。 Unicode是一个字符集，为每个字符规定一个用来表示该字符的数字，但是并没有规定该数字的二进制保存方式，utf8规定了对于unicode值的二进制保存方式
  //大白话就是一组数字对应一个字符
  //string：Go语言中，string就是只读的采用utf8编码的字节切片(slice) 因此用len函数获取到的长度并不是字符个数，而是字节个数。 for循环遍历输出的也是各个字节
  //rune :rune是int32的别名，代表字符的Unicode编码，采用4个字节存储，将string转成rune就意味着任何一个字符都用4个字节来存储其unicode值，这样每次遍历的时候返回的就是unicode值，而不再是字节了，这样就可以解决乱码问题了
  //注意runue不是等于unicode而是以什么形式储存unicode值，我以前就是认为rune代表unicode
  //大白话来讲：go中是utf-8编码 而unicode是字符集就是一数字来映射（代表）字符，这就是utf-8和unicode的区别，utf-8是以三个字节来存中文，
  //而rune类型是四个四节，任何string类型转化为rune 是以四个字节来存4个字节来存储其unicode值 不管哪种字符
  //字符的unicode值决定了字符需要用多少字节表示
  //字符的unicode值决定了字符需要用多少字节表示
  //就是rune类型int32go语言中是代表unicode编码

	s1 :="截11取22中文"
    fmt.Println([]rune(s1))//25130 49 49 21462 50 50 20013 25991] //unicode码点
    //go直接用4个四节的rune类型来代表unicode节点所以想到rune就得想到它是unicode码点

    //为什么不是三个字节跟切片的左闭右开原则有关？然而并没有关系  go就是用四个字节来代表unicode码点
    s3 :="a我爱中国"
	fmt.Println(s3[:]) //a我爱中国 为什么不是打印出[230,234,3434...]
	fmt.Println(s3[:2])//a�
	fmt.Println(s3[:4])//a我
	fmt.Println(s3[1]) //230
	//for range循环打印出来的是rune类型
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
	//[]byte->unicode->rune?

	//sting
	/*
	 type stringStruct struct {
	    str unsafe.Pointer
	    len int
	}
	string其实就是个struct。
	可以看到str其实是个指针，指向某个数组的首地址，另一个字段是len长度。那到这个数组是什么呢？ 在实例化这个stringStruct的时候：
	func gostringnocopy(str *byte) string {
		ss := stringStruct{str: unsafe.Pointer(str), len: findnull(str)}
		s := *(*string)(unsafe.Pointer(&ss))
		return s
	}
	其实就是byte数组 --sring的底层指向一个byte数组
	//一个slice由三部分构成：指针、长度和容量，slice的底层引用了一个数组对象，多个slice之间可以共享同一底层数据，
	//指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。
	*/
	fmt.Println("----------------分割线-------------------")
    s6 :="hello world"
    fmt.Println(s6[0]) //104

   s7:=[]byte(s6)
   s7[0]='c'
   fmt.Println(string(s7))

}
