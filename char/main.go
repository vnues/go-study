package main

import "fmt"

//      rune的作用是将unicode编码表示存入 ，而unicode编码是字符集编码的表示
func lengthOfNonRepeatingSubStr(s string) int {
	// map的作用可以让我们判断有没有存在 还可以存入下标
	maxlength := 0
	start := 0
	// 默认值是0
	lastOccurred := make(map[rune]int)
	// 记录字符最后出现的位置 判断他有没有存在 这个map类型可以检测出来 1231 1的下标0 1下标3
	for i, ch := range []rune(s) {
		// 大于等于的情况下
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			// 说明重复啦 继续查找
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("aback"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bobby"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("poke"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abide"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
	s1 := "hello 世界hello123"
	println(lengthOfNonRepeatingSubStr(s1))
}
