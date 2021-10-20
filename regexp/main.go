package main

import (
	"fmt"
	"regexp"
)

// 全局匹配or局部匹配
/*
正则表达式都是针对于字符串（还有数字这类的）所以string方法会很多
go正则的ALL指的是全局搜索 方法没有all不是全局搜索 all全局搜索如果需要int参数
-1代表返回全部 0表示一个
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
	//代表如果小于0返回全部匹配到的字段
	if n < 0 {
		n = len(s) + 1
	}
	var result [][]string
	re.allMatches(s, nil, n, func(match []int) {
		if result == nil {
			result = make([][]string, 0, startSize)
		}
		slice := make([]string, len(match)/2)
		for j := range slice {
			if match[2*j] >= 0 {
				slice[j] = s[match[2*j]:match[2*j+1]]
			}
		}
		result = append(result, slice)
	})
	return result
}
*/
// 正则字符含义
/*
	代码	说明
	.	匹配除换行符以外的任意字符
	\w	匹配字母或数字或下划线或汉字
	\s	匹配任意的空白符
	\d	匹配数字
	\b	匹配单词的开始或结束
	^	匹配字符串的开始
	$	匹配字符串的结束
*/
// 正则方法
/*
regexp是公共方法 r是结构体
regexp.MatchString
regexp.Compile
r.MatchString
r.FindString
r.FindStringIndex
r.FindStringSubmatch 这个方法返回全局匹配的字符串和局部匹配的字符，比如
FindStringSubmatchIndex 和上面的方法一样，不同的是返回全局匹配和局部匹配的 起始索引和结束索引
r.FindAllString 个方法返回所有正则匹配的字符，不仅仅是第一个（全局搜索/匹配）
r.FindAllStringSubmatchIndex 这个方法返回所有全局匹配和局部匹配的字符串起始索引和结束索引
r.FindAllString  为这个方法提供一个正整数参数来限制匹配数量
regexp.MustCompile  当使用正则表达式来创建常量的时候，你可以使用`MustCompile` 因为`Compile`返回两个值
r.ReplaceAllString
r.ReplaceAllFunc
*/

/*
[里面的表达式是||或者不是&&]
[]//里面放的值就是实际的值 不是正则的匹配符号那样的 比如你a-z表示a-z的范围，放.就表示点 而不是表示任意字符
[a-z.]==>这样理解为 a-z||.(小数点)
compile生成正则表达式
*/
func main() {

	// flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	// params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	// fmt.Println(params)
	// for _,param :=range params {
	//	fmt.Println(param)
	// }
	// FindStringSubmatch方法是提取出匹配的字符串，然后通过[]string返回。我们可以看到，第1个匹配到的是这个字符串本身，从第2个开始，才是我们想要的字符串。

	// 测试模式是否匹配字符串，括号里面的意思是
	// 至少有一个a－z之间的字符存在
	// match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//	fmt.Println(match)
	// 上面我们直接使用了字符串匹配的正则表达式，
	// 但是对于其他的正则匹配任务，你需要使用
	// `Compile`来使用一个优化过的正则对象
	// compile编译
	r, _ := regexp.Compile("p([a-z]+)ch")
	/*
		*Regexp
		 func Compile(expr string) (*Regexp, error) {
			return compile(expr, syntax.Perl, false)
		}
	*/
	// r是正则结构体对象
	// fmt.Println(r) //p([a-z]+)ch
	// fmt.Printf("%T \n",r) //*regexp.Regexp  这样子拿到的确实是值而不是地址
	// 正则结构体对象有很多方法可以使用，比如上面的例子
	// 也可以像下面这么写
	// fmt.Println(r.MatchString("peach"))

	// fmt.Println(r.FindString("11peach11"))
	// 这个方法返回全局匹配的字符串和局部匹配的字符，比如
	// 这里会返回匹配`p([a-z]+)ch`的字符串
	// 和匹配`([a-z]+)`的字符串
	// 局部匹配的就是以这种[]形式出现的局部单元
	// 以数组形式返回
	fmt.Println(r.FindStringSubmatch("peawwwwch  pea11111wwwwch hello world"))
	fmt.Println(r.FindAllStringSubmatch("peawwwwch  pewerwerwerwerdfdsfwfswwch hello world", 2))
	matches := r.FindStringSubmatch("peawwwwch  pea11111wwwwch hello world")
	for _, m := range matches {
		fmt.Println(m)
	}
	allMatches := r.FindAllStringSubmatch("peawwwwch  pewerwerwerwerdfdsfwfswwch hello world", -1)
	for _, all := range allMatches {
		// [][]string
		// string-byte 你可以理解成一维byteslicel类型就是string
		/*
		 func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
		*/
		fmt.Println("--------------------string分割线-------------------------")

		fmt.Println(all)
	}
	allMatchesByte := r.FindAllSubmatch([]byte("peawwwwch  pewerwerwerwerdfdsfwfswwch hello world"), -1)
	fmt.Println(allMatchesByte)
	for _, allByte := range allMatchesByte {
		fmt.Println("--------------------分割线-------------------------")
		// fmt.Println(string(allByte)) string只能转化一维数组的
		fmt.Println(allByte)    // allByte是二维数组 allByte[1]就是它的子项·	去啊在Aq·	A·112awZA1·
		fmt.Println(allByte[1]) // allByte[1]得到的是一个一维数组
		fmt.Println(string(allByte[1]))
	}
	notall := r.FindSubmatch([]byte("peawwwwch  pewerwerwerwerdfdsfwfswwch hello world"))
	fmt.Println(notall)
	for _, not := range notall {
		fmt.Println("--------------------not分割线-------------------------")
		fmt.Println(not)
		fmt.Printf("%s \n", not)
	}
	// 来循环这个数组
	// params :=  r.FindStringSubmatch("peawwwwch hello world")
	// 循环的固定形式吗？第一个是index，第二个是value
	// for _,match:= range params{
	//	fmt.Println(match)
	//
	// }
	/*
		func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
			if n < 0 {
				n = len(s) + 1
			}
			var result [][]string
			re.allMatches(s, nil, n, func(match []int) {
				if result == nil {
					result = make([][]string, 0, startSize)
				}
				slice := make([]string, len(match)/2)
				for j := range slice {
					if match[2*j] >= 0 {
						slice[j] = s[match[2*j]:match[2*j+1]]
					}
				}
				result = append(result, slice)
			})
			return result
		}
	*/
	// n int参数 限制匹配数量
	//	// 这个方法返回所有正则匹配的字符，不仅仅是第一个
	// fmt.Println(r.FindAllStringSubmatch("peaddddchqwqweqwe hello world",-1))
	// fmt.Println(r.FindStringIndex("peaddddcpeaddddchqwqweqwecpeaddddc hello world peaddddc"))
	// 这个方法返回所有全局匹配和局部匹配的字符串起始索引
	// 和结束索引
	// [0 5 1 3] [6 11 7 9] [12 17 13 15]] 返回的是全部的索引以及局部的索引
	// go 正则结束的位置就是你匹配到了以后继续找 发现没有匹配到就结束，这个搜索的位置就是结束位置
	// fmt.Println(r.FindAllStringSubmatchIndex("peach punch punch", 4))
	// 为这个方法提供一个正整数参数来限制匹配数量

	// fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 这个方法查找第一次匹配的索引，并返回匹配字符串
	// 的起始索引和结束索引，而不是匹配的字符串
	// fmt.Println(r.FindStringIndex("peach punch"))

	/*

		[]byte byte数组
		Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。
		这两个名称可以互换使用。同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。
	*/

}
